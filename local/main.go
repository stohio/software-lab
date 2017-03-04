package main

import (
	"log"
	"net"
	"strings"
	"os"
	//"net/http"
)

func main() {
	log.Printf("Starting Local Server...")
	localIP := GetOutboundIP()
	log.Printf("Local IP: %s", localIP)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Hostname: %s", hostname)

	node := []byte(`{
		"name": ` + hostname + `,
		"local_ip": ` + $localIP + `
	}`)



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
