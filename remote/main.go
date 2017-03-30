package main

import (
	"log"
	"net/http"
)

// main creates a new router instance and starts running the http server on port 8080
func main() {
	router := NewRouter()

	log.Printf("Remote Server is Running")
	log.Fatal(http.ListenAndServe(":8080", router))
}
