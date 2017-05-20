package main

import (
	"net/http"

	swl "github.com/stohio/software-lab/lib"
)

var remotePort = swl.GetRemotePort()

// main creates a new router instance and starts running the http server on port 8080
func main() {

	router := swl.NewRouter(routes)
	swl.InitLogger()

	swl.ConsoleLog.Printf("Remote Server is Running")
	swl.ConsoleLog.Fatal(http.ListenAndServe(":"+remotePort, router))
}
