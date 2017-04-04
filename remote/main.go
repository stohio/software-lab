package main

import (
	"net/http"

	swl "github.com/stohio/software-lab/lib"
)

// main creates a new router instance and starts running the http server on port 8080
func main() {
	router := swl.NewRouter(routes)
	swl.InitLogger()

	swl.ConsoleLog.Printf("Remote Server is Running")
	swl.ConsoleLog.Fatal(http.ListenAndServe(":8080", router))
}
