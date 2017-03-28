package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	swl "github.com/stohio/software-lab/lib"
)

// Test prints a confirmation that the server is up and accepting requests
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up")
}

// NetworkIndex currently doesn't do anything, and I'm not sure what it should be doing
// TODO
func NetworkIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(networks); err != nil {
		panic(err)
	}
}

// NetworkCreate takes in a network's name, ip, and stack id and creates a new local server based on this information
func NetworkCreate(w http.ResponseWriter, r *http.Request) {
	var netCreate swl.NetworkCreate

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if !ValidateJSON(body, &netCreate, w) {
		return
	}

	if !ValidateParam("name", netCreate.Name, w) {
		return
	}

	if !ValidateParamRegex("ip", netCreate.IP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w) {
		return
	}

	if !ValidateParamRegex("stack", netCreate.Stack, "\\A[\\d]+\\z", w) {
		return
	}

	stack := RepoFindStack(*netCreate.Stack)

	if stack == nil {
		response := ParamError{
			Error: "Stack Not Found",
			Param: "stack",
			Value: string(*netCreate.Stack),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(409)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
		return
	}

	// create the first node of this network, based off the requesting machine
	// TODO should IP be the GetIPAddress(r) or the ip of the network being created?
	node := swl.Node{
		Name: netCreate.Name,
		IP:   netCreate.IP,
	}
	n := RepoCreateNode(&node)

	// TODO n is appended to nodes in RepoCreateNode
	// var nodes swl.Nodes
	// nodes = append(nodes, n)
	netAddr := GetIPAddress(r)

	// TODO alright this doesn't make any sense. WTF
	// First shouldn't network be created before the node
	// Second shouldn't the attributes of network be based off the body of the request
	// nodes is the list of every node for every network right?
	network := swl.Network{
		IP:    netAddr,
		Nodes: nodes,
		Stack: stack,
	}

	net := RepoCreateNetwork(&network)

	w.Header().Set("Content-Type", "application/json; charset=UTH-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(net); err != nil {
		panic(err)
	}
}

func NodeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		panic(err)
	}
}

func NodeCreate(w http.ResponseWriter, r *http.Request) {
	var node swl.Node
	//Get JSON string, ensure it isn't too big
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	//Close handler
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if !ValidateJSON(body, &node, w) {
		return
	}

	if !ValidateParam("name", node.Name, w) {
		return
	}

	if !ValidateParamRegex("ip", node.IP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w) {
		return
	}

	fmt.Println("ABOUT TO NETADDR")

	netAddr := GetIPAddress(r)
	lastDot := -1
	for i, c := range netAddr {
		if c == '.' {
			lastDot = i
		}
	}
	fmt.Printf("NET ADDR %d\n", lastDot)
	fmt.Println(netAddr[:lastDot])

	if network := RepoFindNetworkByIP(netAddr); network != nil {
		n := RepoCreateNode(&node)
		network.Nodes = append(network.Nodes, n)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(network); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(409)
		if err := json.NewEncoder(w).Encode(stacks); err != nil {
			panic(err)
		}
	}
}

func NodeEnable(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nodeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	if n := RepoEnableNode(nodeID); n != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(200)
		if err := json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	} else {
		paramError := ParamError{
			Error: "Node with id not found",
			Param: "id",
			Value: vars["id"],
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(406)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}
}

func NodeGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nodeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	if n := RepoFindNode(nodeID); n != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(200)
		if err := json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	} else {
		paramError := ParamError{
			Error: "Node with id not found",
			Param: "id",
			Value: strconv.Itoa(nodeID),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(406)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}
}

func NodeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	if err := DeleteNode(nodeID); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(406)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}

func NodeIncrementClients(w http.ResponseWriter, r *http.Request) {
	NodeUpdateClients(w, r, true)
}

func NodeDecrementClients(w http.ResponseWriter, r *http.Request) {
	NodeUpdateClients(w, r, false)
}

func NodeUpdateClients(w http.ResponseWriter, r *http.Request, increment bool) {
	vars := mux.Vars(r)
	nodeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	if err := RepoUpdateNodeClients(nodeID, increment); err != nil {
		paramError := ParamError{
			Error: "Nodewith id not found",
			Param: "id",
			Value: strconv.Itoa(nodeID),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(406)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
}

func NetworkCurrent(w http.ResponseWriter, r *http.Request) {
	netAddr := GetIPAddress(r)
	net := RepoFindNetworkByIP(netAddr)
	if net == nil {
		paramError := ParamError{
			Error: "No Network Found",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(net); err != nil {
		panic(err)
	}
}

func NetworkGetNodeDownload(w http.ResponseWriter, r *http.Request) {
	netAddr := GetIPAddress(r)
	node := RepoFindBestNodeInNetworkByIP(netAddr)
	if node == nil {
		paramError := ParamError{
			Error: "Could Not Find a Node",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}

	//otherwise we need to make a requst to the node to get the software
	//we should send to the node who made the request.
	//And thenode should return the software to taht IP address I guess
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(node); err != nil {
		panic(err)
	}
}

func SoftwareGet(w http.ResponseWriter, r *http.Request) {
	NetworkGetNodeDownload(w, r)
}

func PackageGet(w http.ResponseWriter, r *http.Request) {
	NetworkGetNodeDownload(w, r)
}
