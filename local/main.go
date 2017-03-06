package main

import (
	"log"
	"net"
	"strings"
	"os"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"bytes"

	swl "github.com/stohio/software-lab/lib"

)

var remoteURL string
var network swl.Network
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
}


func DownloadSoftware(initial bool) {
	if initial {
		fmt.Println("Inital Download")
	} else {
		fmt.Println("Additional Download")
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
