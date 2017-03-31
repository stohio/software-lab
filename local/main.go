package main

import (
	"encoding/json"
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

	"github.com/franela/goreq"
	swl "github.com/stohio/software-lab/lib"
)

const remoteURL = "http://stoh.io/swl"

var network swl.Network
var node swl.Node

var client *http.Client

func main() {
	log.Printf("Starting Local Server...")
	localIP := GetOutboundIP()
	log.Printf("Local IP: %s", localIP)
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
			IP:    &localIP,
			Name:  &hostname,
			Stack: &stackID,
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

	//fmt.Printf("%d\n", node.Id)
	// Enable the node
	EnableNode()

	//Now it needs to serve its routes
	router := NewRouter()

	log.Printf("The Node is now ready to serve files!")
	log.Fatal(http.ListenAndServe(":80", router))
}

//EnableNode sends the enable POST
func EnableNode() {

	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.Id) + "/enable",
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

//DeleteNode sends the DELETE post
func DeleteNode() {
	resp, err := goreq.Request{
		Method: "DELETE",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.Id),
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

//AddClient sends a POST to increment
func AddClient() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.Id) + "/clients/increment",
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

//RemoveClient sends a POST to decrement
func RemoveClient() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri:    remoteURL + "/nodes/" + strconv.Itoa(node.Id) + "/clients/decrement",
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

//DownloadSoftware will download the software for rehosting
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
		path := "software/" + strconv.Itoa(s.Id)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
		for _, v := range s.Versions {
			filename := path + "/" + strconv.Itoa(v.Id) + v.Extension
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
						Uri:    remoteURL + "/software/" + strconv.Itoa(s.Id) + "/versions/" + strconv.Itoa(v.Id),
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
						Uri:    "http://" + *node.IP + "/download/software/" + strconv.Itoa(s.Id) + "/versions/" + strconv.Itoa(v.Id),
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

//SetupInitialNode runs through process to select a stack
func SetupInitialNode(stacks swl.Stacks) int {
	for _, s := range stacks {
		fmt.Printf("(%d) - %s\n", s.Id, s.Name)
	}
	var response int
	if _, err := fmt.Scanf("%d", &response); err != nil {
		fmt.Println("Invalid Response")
		return SetupInitialNode(stacks)
	}
	for _, s := range stacks {
		if s.Id == response {
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

func GetNode() swl.Node {
	return node
}

//type Node struct {
//	Id	int		`json:"id"`
//	Name	*string		`json:"name"`
//	IP	*string		`json:"ip"`
//	Enabled bool		`json:"enabled"`
//	Clients	int		`json:"clients"`
//	Added	time.Time	`json:"added"`
//}
