package main

import (
	"net/http"
)

// main creates a new router instance and starts running the http server on port 8080
func main() {
	router := NewRouter()
	InitLogger()

	consoleLog.Printf("Remote Server is Running")
	consoleLog.Fatal(http.ListenAndServe(":8080", router))
}
