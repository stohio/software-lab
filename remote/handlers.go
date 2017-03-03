package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"

)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up")
}

func NodeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		panic(err)
	}
}

func NodeCreate(w http.ResponseWriter, r *http.Request) {
	var node Node
	//Get JSON string, ensure it isn't too big
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	//Close handler
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if !ValidateJson(body, &node, w) {
		return
	}

	if !ValidateParam("name", node.Name, w) {
		return
	}

	if !ValidateParamRegex("local_ip", node.LocalIP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w){
		return
	}

	n := RepoCreateNode(node)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(n); err != nil {
		panic(err)
	}
}
