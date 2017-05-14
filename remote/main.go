package main

import (
	"log"
	"net"
	"net/http"
	"strings"

	swl "github.com/stohio/software-lab/lib"
)

// main creates a new router instance and starts running the http server on port 8080
func main() {
	router := swl.NewRouter(routes)
	swl.InitLogger()

	localIP := GetOutboundIP()

	swl.ConsoleLog.Printf("Remote Server is Running %s:8080", localIP)
	swl.ConsoleLog.Fatal(http.ListenAndServe(":8080", router))
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
