package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Printf("Starting Local Server...\n")
	fmt.Printf("Local Address: " + GetOutboundIP() + "\n")
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}
