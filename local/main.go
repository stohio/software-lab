package main

import (
	"log"
	"net"
	"strings"
	"os"
	"encoding/json"
	"io"
	"io/ioutil"
	"fmt"
	"net/http"
	"strconv"

	"github.com/franela/goreq"
	swl "github.com/stohio/software-lab/lib"

)

//const remoteURL = "https://stoh.io/swl"
const remoteURL = "http://127.0.0.1:8080"
var network swl.Network
var node swl.Node

var client *http.Client

func main() {
	log.Printf("Starting Local Server...")
	localIP := GetOutboundIP()
	log.Printf("Local IP: %s", localIP)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Hostname: %s", hostname)
	node = swl.Node{
		Name:	&hostname,
		IP:	&localIP,
	}


	resp, err := goreq.Request{
		Method: "POST",
		Uri: remoteURL + "/nodes",
		Body: node,
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

		newNet := swl.NetworkCreate {
			IP:	&localIP,
			Name:	&hostname,
			Stack:	&stackID,
		}
                //Funyction called SimpleRequest that takes the object to be JSONified
                //and returns the oh maybe this wont work since we need to close the res
		//req, err = http.NewRequest("POST", remoteURL + "/networks", bytes.NewBuffer(jsonBytes))
		//req.Header.Set("Content-Type", "application/json")
		//resp, err = client.Do(req)

		resp, err = goreq.Request{
			Method: "POST",
			Uri: remoteURL + "/networks",
			Body: newNet,
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

        //If there is a network or existing node already setup dont download from the internet
        //we will instead copy from an existing node
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

        log.Printf("The Node is now ready to serve files!");
        log.Fatal(http.ListenAndServe(":5000", router))
}

func EnableNode() {

    resp, err := goreq.Request{
	    Method: "POST",
	    Uri: remoteURL + "/nodes/" + strconv.Itoa(node.Id) + "/enable",
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
func AddClient() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri: remoteURL + "/nodes/" + strconv.Itoa(node.Id) + "/clients/increment",
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

func DownloadSoftware(initial bool) {
	if _, err := os.Stat("software"); os.IsNotExist(err) {
		os.Mkdir("software", 0755)
	}
	CheckOrDownload(network.Stack.Softwares, initial)
	for _, p := range network.Stack.Packages {
		CheckOrDownload(p.Softwares, initial)
	}
}

func CheckOrDownload(softwares swl.Softwares,initial bool) {
	for _, s := range softwares {
		path := "software/" + strconv.Itoa(s.Id)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
		for _, v := range s.Versions {
			filename := path + "/" + strconv.Itoa(v.Id) + v.Extension
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				if initial {
					fmt.Printf("Downloading %s ...\n", filename)
					out, err := os.Create(filename)
					if err != nil {
                                                fmt.Println("here is another panic")
						panic(err)
					}
					defer out.Close()
					resp, err := http.Get(v.URL)
					if err != nil {
                                                fmt.Println("and another here is another panic")
						panic(err)
					}
					_, err = io.Copy(out, resp.Body)
					if err != nil {
                                                fmt.Println("and theres one more! here is another panic")
						panic(err)
					}
					fmt.Printf("Downloaded %s\n", filename)
				} else {
                                        //Make some call to the current node to download
                                        //Get node
                                        //Make the request to get ALLLL the times
                                        //We can use that "." cam found
					fmt.Printf("Need to Download %s locally\n", filename)
				}
			}
		}
	}
}


func SetupInitialNode(stacks swl.Stacks) int {
	for _, s := range stacks {
		fmt.Printf("(%d) - %s\n", s.Id, s.Name)
	}
	var response int
	if _, err := fmt.Scanf("%d",&response); err != nil {
		fmt.Println("Invalid Response\n")
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
