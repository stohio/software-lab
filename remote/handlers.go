package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up")
}

func NodesIndex(w http.ResponseWriter, r *http.Request) {
	nodes := Nodes {
		Node {Name: "Node 1", LocalIP: "10.0.0.20", Network: "100.10.20.30"},
		Node {Name: "Noded 2", LocalIP: "10.0.0.35", Network: "100.10.20.30"},
	}

	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		panic(err)
	}
}
