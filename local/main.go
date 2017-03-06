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
	"bytes"
	"strconv"

	swl "github.com/stohio/software-lab/lib"

)

var remoteURL string
var network swl.Network
var node swl.Node
func main() {

	remoteURL = "http://127.0.0.1:8080"
	log.Printf("Starting Local Server...")
	localIP := GetOutboundIP()
	log.Printf("Local IP: %s", localIP)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Hostname: %s", hostname)
	node := swl.Node{
		Name:	&hostname,
		IP:	&localIP,
	}

	jsonBytes, _ := json.Marshal(node)

	req, err := http.NewRequest("POST", remoteURL + "/nodes", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 409 {
		var stacks swl.Stacks
		if err := json.Unmarshal(body, &stacks); err != nil {
			panic(err)
		}
		fmt.Println("This Node is the initial node.  Please choose a stack to use.")
		stackID := SetupInitialNode(stacks)
		fmt.Printf("Stack %d was selected\n", stackID)

		newNet := swl.NetworkCreate {
			IP:	&localIP,
			Name:	&hostname,
			Stack:	&stackID,
		}
		jsonBytes, _ = json.Marshal(newNet)

		req, err = http.NewRequest("POST", remoteURL + "/networks", bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ = ioutil.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &network); err != nil {
			panic(err)
		}
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
		}
	}

	jsonBytes = []byte(`{"id": ` + strconv.Itoa(node.Id) + `}`)
	req, err = http.NewRequest("POST", remoteURL + "/nodes/enable", bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Println("Node is Active")
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
