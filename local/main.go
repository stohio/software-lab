package main

import (
	"encoding/json"
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
	"time"

	astilectron "github.com/asticode/go-astilectron"
	astilog "github.com/asticode/go-astilog"
	"github.com/franela/goreq"
	"github.com/pkg/errors"
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

	//Init GUI

	var a *astilectron.Astilectron
	var err error
	if a, err = astilectron.New(astilectron.Options{
		AppName:           "Software Lab",
		BaseDirectoryPath: "",
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "creating new astilectron failed"))
	}
	defer a.Close()
	a.HandleSignals()

	if err := a.Start(); err != nil {
		astilog.Fatal(errors.Wrap(err, "starting failed"))
	}

	var w *astilectron.Window
	if w, err = a.NewWindow("ui/index.html", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "new window failed"))
	}
	if err := w.Create(); err != nil {
		astilog.Fatal(errors.Wrap(err, "creating window failed"))
	}

	a.Wait()

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

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		DeleteNode()
		os.Exit(1)
	}()

	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes",
		Body:   node,
	}.Do()
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//Returns body

	//If there are no existing nodes, create a network with a user-defined stack
	if resp.StatusCode == 409 {
		//Maybe wrap this all in a nice function called CreateNetwork
		var stacks swl.Stacks
		if err := json.Unmarshal(body, &stacks); err != nil {
			panic(err)
		}
		fmt.Println("This Node is the initial node.  Please choose a stack to use.")
		stackID := SetupInitialNode(stacks)
		fmt.Printf("Stack %d was selected\n", stackID)
		//done with stack code

		newNet := swl.NetworkCreate{
			IP:      &localIP,
			Name:    &hostname,
			StackID: &stackID,
		}
		//Funyction called SimpleRequest that takes the object to be JSONified
		//and returns the oh maybe this wont work since we need to close the res
		//req, err = http.NewRequest("POST", remoteURL + "/networks", bytes.NewBuffer(jsonBytes))
		//req.Header.Set("Content-Type", "application/json")
		//resp, err = client.Do(req)

		resp, err = goreq.Request{
			Method: "POST",
			Uri:    remoteURL + "/networks",
			Body:   newNet,
		}.Do()

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ = ioutil.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &network); err != nil {
			panic(err)
		}
		//returns body
		DownloadSoftware(true)

	} else if resp.StatusCode == 201 {
		if err := json.Unmarshal(body, &network); err != nil {
			panic(err)
		}
		DownloadSoftware(false)

	} else {
		panic("Unexpected Response Code")
	}

	for _, n := range network.Nodes {
		if *n.IP == localIP {
			node = *n
			break
		}
	}

	//fmt.Printf("%d\n", node.ID)
	// Enable the node
	EnableNode()

	//Now it needs to serve its routes
	router := swl.NewRouter(routes)

	log.Printf("The Node is now ready to serve files!")
	log.Fatal(http.ListenAndServe(":"+localPort, router))
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

//CheckOrDownload will check to see if software needs downloaded
func CheckOrDownload(softwares swl.Softwares, initial bool) {
	for _, s := range softwares {
		path := "software/" + strconv.Itoa(s.ID)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
		for _, v := range s.Versions {
			filename := path + "/" + strconv.Itoa(v.ID) + v.Extension
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				time.Sleep(time.Second * 2)
				if initial {
					fmt.Printf("Downloading %s - %s ...\n", s.Name, v.OS)
					out, err := os.Create(filename)
					if err != nil {
						panic(err)
					}
					defer out.Close()
					resp, err := http.Get(v.URL)
					if err != nil {
						panic(err)
					}
					_, err = io.Copy(out, resp.Body)
					if err != nil {
						panic(err)
					}
					fmt.Printf("Downloaded %s\n", filename)
				} else {
					out, err := os.Create(filename)
					defer out.Close()
					resp, err := goreq.Request{
						Method: "GET",
						Uri:    remoteURL + "/software/" + strconv.Itoa(s.ID) + "/versions/" + strconv.Itoa(v.ID),
					}.Do()
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()
					body, _ := ioutil.ReadAll(resp.Body)
					var node swl.Node
					if err := json.Unmarshal(body, &node); err != nil {
						fmt.Println(string(body))
						panic(err)
					}

					resp, err = goreq.Request{
						Method: "GET",
						Uri:    "http://" + *node.IP + "/download/software/" + strconv.Itoa(s.ID) + "/versions/" + strconv.Itoa(v.ID),
					}.Do()
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()
					_, err = io.Copy(out, resp.Body)
					if err != nil {
						panic(err)
					}

					fmt.Printf("Copied the file %s\n", filename)
				}
			}
		}
	}
}

//SetupInitialNode prints all the stacks available and accepts user input picking one of the stacks for this network
func SetupInitialNode(stacks swl.Stacks) int {
	for _, s := range stacks {
		fmt.Printf("(%d) - %s\n", s.ID, s.Name)
	}
	var response int
	if _, err := fmt.Scanf("%d", &response); err != nil {
		fmt.Println("Invalid Response")
		return SetupInitialNode(stacks)
	}
	for _, s := range stacks {
		if s.ID == response {
			return response
		}
	}
	fmt.Printf("%d is not a Stack\n\n", response)
	return SetupInitialNode(stacks)
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
