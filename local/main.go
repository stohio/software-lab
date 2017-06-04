package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/franela/goreq"
	swl "github.com/stohio/software-lab/lib"
)

// const remoteURL = "http://stoh.io/swl"

var remoteServer = swl.GetRemoteServer()
var remotePort = swl.GetRemotePort()
var defaultRemoteURL = remoteServer + ":" + remotePort

var localPort = swl.GetLocalPort()

var remoteURL string
var client *http.Client

func main() {

	//Flags and etc. for initilization
	remotePtr := flag.String("remote", defaultRemoteURL, "IP Address of remote server")
	flag.Parse()
	remoteURL = "http://" + *remotePtr

	log.Printf("Starting Local Server -> %s", *remotePtr)
	localIP := GetOutboundIP()
	log.Printf("Local IP: %s", localIP)
	swl.InitLogger()
	hostname, err := os.Hostname()
	if err != nil {
		swl.ConsoleLog.Fatal(err)
	}
	log.Printf("Hostname: %s", hostname)
	node = swl.Node{
		Name: &hostname,
		IP:   &localIP,
	}

	//Allows smooth exit if needed.
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		DeleteNode()
		os.Exit(1)
	}()

	//Initialize node
	initNode()

	//TODO REMOVE

	//Old Code
	//	resp, err := goreq.Request{
	//		Method: "POST",
	//		Uri:    remoteURL + "/nodes",
	//		Body:   node,
	//	}.Do()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	defer resp.Body.Close()
	//
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	//Returns body
	//
	//	//If there are no existing nodes, create a network with a user-defined stack
	//	if resp.StatusCode == 409 {
	//		//Maybe wrap this all in a nice function called CreateNetwork
	//		var stacks swl.Stacks
	//		if err := json.Unmarshal(body, &stacks); err != nil {
	//			panic(err)
	//		}
	//		fmt.Println("This Node is the initial node.  Please choose a stack to use.")
	//		stackID := selectStackCLI(stacks)
	//		fmt.Printf("Stack %d was selected\n", stackID)
	//		//done with stack code
	//
	//		newNet := swl.NetworkCreate{
	//			IP:      &localIP,
	//			Name:    &hostname,
	//			StackID: &stackID,
	//		}
	//		//Funyction called SimpleRequest that takes the object to be JSONified
	//		//and returns the oh maybe this wont work since we need to close the res
	//		//req, err = http.NewRequest("POST", remoteURL + "/networks", bytes.NewBuffer(jsonBytes))
	//		//req.Header.Set("Content-Type", "application/json")
	//		//resp, err = client.Do(req)
	//
	//		resp, err = goreq.Request{
	//			Method: "POST",
	//			Uri:    remoteURL + "/networks",
	//			Body:   newNet,
	//		}.Do()
	//
	//		if err != nil {
	//			panic(err)
	//		}
	//		defer resp.Body.Close()
	//		body, _ = ioutil.ReadAll(resp.Body)
	//
	//		if err := json.Unmarshal(body, &network); err != nil {
	//			panic(err)
	//		}
	//		//returns body
	//		DownloadSoftware(true)
	//
	//	} else if resp.StatusCode == 201 {
	//		if err := json.Unmarshal(body, &network); err != nil {
	//			panic(err)
	//		}
	//		DownloadSoftware(false)
	//
	//	} else {
	//		panic("Unexpected Response Code")
	//	}
	//
	//	for _, n := range network.Nodes {
	//		if *n.IP == localIP {
	//			node = *n
	//			break
	//		}
	//	}
	//
	//	//fmt.Printf("%d\n", node.ID)
	//	// Enable the node
	//	EnableNode()
	//
	//Now it needs to serve its routes
	router := swl.NewRouter(routes)

	log.Printf("The Node is now ready to serve files!")
	log.Fatal(http.ListenAndServe(":"+localPort, router))
}

func initNode() {

	//Check if a network exists already
	resp, err := goreq.Request{
		Method: "GET",
		Uri:    remoteURL + "/networks/current",
	}.Do()

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//If the network doesn't exist yet, create the network
	if resp.StatusCode == 404 {
		createNetwork()
	} else if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &network); err != nil {
			panic(err)
		}
	}

	//Add node to the network
	resp, err = goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes",
		Body:   node,
	}.Do()

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		panic("Unexpected Response Code From POST /nodes.  " +
			"Expected 201, Received " + strconv.Itoa(resp.StatusCode))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &node); err != nil {
		panic(err)
	}

	setupSoftware()
	EnableNode()

}

func createNetwork() {
	//Get available stacks
	resp, err := goreq.Request{
		Method: "GET",
		Uri:    remoteURL + "/stacks",
	}.Do()

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//Something went wrong!
		panic("Unexpected Response Code From /stacks. " +
			"Expected 200, Received " + strconv.Itoa(resp.StatusCode))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var stacks swl.Stacks
	if err := json.Unmarshal(body, &stacks); err != nil {
		panic(err)
	}

	//Let user select stack
	fmt.Println("This Node is the initial node.  Please choose a stack to use.")
	stackID := selectStackCLI(stacks)
	fmt.Printf("Stack %d was selected\n", stackID)

	newNet := swl.NetworkCreate{
		StackID: &stackID,
	}

	resp, err = goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/networks",
		Body:   newNet,
	}.Do()

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		panic("Unexpected Response Code From POST /networks.  " +
			"Expected 201, Received " + strconv.Itoa(resp.StatusCode))
	}

	body, _ = ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &network); err != nil {
		panic(err)
	}

}

//EnableNode sends a request to the remote server setting this local server to enabled
func EnableNode() {

	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.ID) + "/enable",
	}.Do()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Println("Node is Active")
	} else {
		fmt.Println(string(body))
	}

}

//DeleteNode sends a request to the remote server deleting this local server from the list of nodes
func DeleteNode() {
	resp, err := goreq.Request{
		Method: "DELETE",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.ID),
	}.Do()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Node is Deleted")
	} else {
		fmt.Println("Something Went Wrong")
		fmt.Println(resp.StatusCode)
	}
}

//AddClient sends a request to the remote server to increment the number of clients for this local server
func AddClient() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.ID) + "/clients/increment",
	}.Do()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		fmt.Println("Node Incremented Clients")
	} else {
		fmt.Println(string(body))
	}
}

//RemoveClient sends a request to the remote server to decrement the number of clients for this local server
func RemoveClient() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.ID) + "/clients/decrement",
	}.Do()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		fmt.Println("Node Decremented Clients")
	} else {
		fmt.Println(string(body))
	}
}

// DownloadSoftware downlaods the associated stack of software to the software directory
// if the software directory doesn't exist it is created
func DownloadSoftware(initial bool) {
	if _, err := os.Stat("software"); os.IsNotExist(err) {
		os.Mkdir("software", 0755)
	}
	CheckOrDownload(network.Stack.Softwares, initial)
	for _, p := range network.Stack.Packages {
		CheckOrDownload(p.Softwares, initial)
	}
}

func setupSoftware() {
	//Create software directory
	os.MkdirAll("softwarelab/software", 0777)

	//Check all software in the stack
	checkSoftware(network.Stack.Softwares)
	//Check all software in Packages
	for _, p := range network.Stack.Packages {
		checkSoftware(p.Softwares)
	}
}

func checkSoftware(softwares swl.Softwares) {
	for _, s := range network.Stack.Softwares {
		path := "softwarelab/software/" + strconv.Itoa(s.ID)
		os.MkdirAll(path, 0777)
		for _, v := range s.Versions {
			//Check if the file has already been downloaded.
			//If it hasn't, download the software
			filename := path + "/" + strconv.Itoa(v.ID) + v.Extension
			if !checkVersionFileIntegrity(filename, v) {
				source := getSoftwareDownloadSource(s, v)
				fmt.Printf("Downloading %s - %s ...\n", s.Name, v.OS)
				downloadSoftware(source, filename, v.Checksum)
			}
		}
	}
}

func getSoftwareDownloadSource(software *swl.Software, version *swl.Version) string {

	//Attempt to request a node to download software from
	resp, err := goreq.Request{
		Method: "GET",
		Uri:    remoteURL + "/software/" + strconv.Itoa(software.ID) + "/versions/" + strconv.Itoa(version.ID),
	}.Do()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//If a node is available, download the software from the node
	if resp.StatusCode == 200 {
		var serverNode swl.Node
		if err := json.Unmarshal(body, &serverNode); err != nil {
			fmt.Println(string(body))
			panic(err)
		}
		return "http://" + *serverNode.IP + "/download/software/" +
			strconv.Itoa(software.ID) + "/versions/" + strconv.Itoa(version.ID)

	} else if resp.StatusCode == 405 {
		return version.URL

	}

	panic("Unexpected Response Code.  Expected 200, 405, " +
		"Received " + strconv.Itoa(resp.StatusCode))

}

func downloadSoftware(source string, path string, checksum string) {
	for i := 0; i < swl.GetLocalDownloadRetries(); i++ {
		resp, err := goreq.Request{
			Method: "GET",
			Uri:    source,
		}.Do()
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var body bytes.Buffer
		bodyTeeReader := io.TeeReader(resp.Body, &body)

		// Check the hash of the body to see if it downloaded correctly
		hash, err := swl.HashFileMd5(bodyTeeReader)
		if err != nil {
			panic(err)
		}
		if hash != checksum {
			fmt.Printf("There was a problem downloading the file: %s"+
				"\nChecksum: %s does not match the download: %s\n",
				path, checksum, hash)
		} else {
			fmt.Printf("Downloaded %s\n", path)
			copyResponseBodyToFile(&body, path)
			return
		}
	}
	panic(errors.New("Software couldn't be downloaded"))

}

//TODO REMOVE

//CheckOrDownload will check to see if software needs downloaded
func CheckOrDownload(softwares swl.Softwares, initial bool) {
	for _, s := range softwares {
		path := "software/" + strconv.Itoa(s.ID)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
		for _, v := range s.Versions {
			filename := path + "/" + strconv.Itoa(v.ID) + v.Extension

			if !checkVersionFileIntegrity(filename, v) {
				if initial {
					downloadFromRemote(s, v, filename)
				} else {
					downloadFromLocal(s, v, filename)
				}
			}
		}
	}
}

// checkFileIntegrity checks to see if a certain version file is downloaded, and if
// it is checks the integrity of it to see if it is a good download. If it is a bad
// download, then it removes the file and returnes false
func checkVersionFileIntegrity(filepath string, version *swl.Version) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	// The file exists
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Check the hash
	fileHash, err := swl.HashFileMd5(file)
	if err != nil {
		panic(err)
	}
	// If hash is bad mark versionDownloadBad as true
	if fileHash != version.Checksum {
		fmt.Printf("%s does not match the checksum, it will be downloaded again"+
			"\nChecksum: %s does not match the download: %s\n",
			filepath, version.Checksum, fileHash)
		// delete this file
		if err = os.Remove(filepath); err != nil {
			panic(err)
		}
		return false
	}

	return true
}

//TODO REMOVE
// downloadFromRemote will download the specified version of software from the remote server
// and save to the file specified by path
func downloadFromRemote(software *swl.Software, version *swl.Version, path string) {
	fmt.Printf("Downloading %s - %s ...\n", software.Name, version.OS)
	for i := 0; i < swl.GetLocalDownloadRetries(); i++ {
		resp, err := goreq.Request{
			Method: "GET",
			Uri:    version.URL,
		}.Do()
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var body bytes.Buffer
		bodyTeeReader := io.TeeReader(resp.Body, &body)

		// Check the hash of the body to see if it downloaded correctly
		hash, err := swl.HashFileMd5(bodyTeeReader)
		if err != nil {
			panic(err)
		}
		if hash != version.Checksum {
			fmt.Printf("There was a problem downloading the file: %s"+
				"\nChecksum: %s does not match the download: %s\n",
				path, version.Checksum, hash)
		} else {
			fmt.Printf("Downloaded %s\n", path)
			copyResponseBodyToFile(&body, path)
			return
		}
	}
	panic(errors.New("Software couldn't be downloaded"))
}

//TODO REMOVE
// downloadFromLocal will download the specified version of software from a local server node
// and save to the file specified by path
func downloadFromLocal(software *swl.Software, version *swl.Version, path string) {
	// query remote for the ip of a node to download the software from
	resp, err := goreq.Request{
		Method: "GET",
		Uri:    remoteURL + "/software/" + strconv.Itoa(software.ID) + "/versions/" + strconv.Itoa(version.ID),
	}.Do()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var node swl.Node
	// unmarshal the response data into a node struct
	if err := json.Unmarshal(body, &node); err != nil {
		fmt.Println(string(body))
		panic(err)
	}

	fmt.Printf("Downloading %s - %s ...\n", software.Name, version.OS)
	for i := 0; i < swl.GetLocalDownloadRetries(); i++ {
		// query the node remote gave us for the software
		resp, err = goreq.Request{
			Method: "GET",
			Uri:    "http://" + *node.IP + "/download/software/" + strconv.Itoa(software.ID) + "/versions/" + strconv.Itoa(version.ID),
		}.Do()
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var body bytes.Buffer
		bodyTeeReader := io.TeeReader(resp.Body, &body)

		// Check the hash of the body to see if it downloaded correctly
		hash, err := swl.HashFileMd5(bodyTeeReader)
		if err != nil {
			panic(err)
		}
		if hash != version.Checksum {
			fmt.Printf("There was a problem downloading the file: %s"+
				"\nChecksum: %s does not match the download: %s\n",
				path, version.Checksum, hash)
		} else {
			fmt.Printf("Copied the file %s\n", path)
			copyResponseBodyToFile(&body, path)
			return
		}
	}

	panic(errors.New("Software couldn't be downloaded"))
}

// copyResponseBodyToFile takes a response containing a piece of software and copies it to
// the file specified by path
func copyResponseBodyToFile(body io.Reader, path string) {
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, body)
	if err != nil {
		panic(err)
	}
}

func selectStackCLI(stacks swl.Stacks) int {
	for _, s := range stacks {
		fmt.Printf("(%d) - %s\n", s.ID, s.Name)
	}
	var response int
	if _, err := fmt.Scanf("%d", &response); err != nil {
		fmt.Println("Invalid Response")
		return selectStackCLI(stacks)
	}
	for _, s := range stacks {
		if s.ID == response {
			return response
		}
	}
	fmt.Printf("%d is not a Stack\n\n", response)
	return selectStackCLI(stacks)
}

//GetOutboundIP dials stohio to get IP address
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "stoh.io:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}

// GetNode returns the node for this machine
func GetNode() swl.Node {
	return node
}

//type Node struct {
//	ID	int		`json:"id"`
//	Name	*string		`json:"name"`
//	IP	*string		`json:"ip"`
//	Enabled bool		`json:"enabled"`
//	Clients	int		`json:"clients"`
//	Added	time.Time	`json:"added"`
//}
