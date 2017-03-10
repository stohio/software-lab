package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Printf("Remote Server is Running")
	log.Fatal(http.ListenAndServe(":80", router))
}
