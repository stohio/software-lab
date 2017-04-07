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

// NetworkIndex responds with networks from repo.go encoded into a JSON array
// @param w: response writer sends a JSON array including all the networks from repo.go
func NetworkIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(networks); err != nil {
		panic(err)
	}
}

// NetworkCreate creates a new network using the requests ip address, and the local ip and name of the initial node as specified in the body.
// @param w: the request should contain the name of the first node, the internal ip of the first node, and the stack id
// of the stack for this network
func NetworkCreate(w http.ResponseWriter, r *http.Request) {
	var netCreate swl.NetworkCreate

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if !swl.ValidateAndUnmarshalJSON(body, &netCreate, w) {
		return
	}

	if !swl.ValidateParam("name", netCreate.Name, w) {
		return
	}

	if !swl.ValidateParamRegex("ip", netCreate.IP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w) {
		return
	}

	stack := RepoFindStack(*netCreate.Stack)

	if stack == nil {
		response := swl.ParamError{
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

	// create the first node of this network, based off the name and IP gotten in the request
	node := swl.Node{
		Name: netCreate.Name,
		IP:   netCreate.IP,
	}
	n := RepoCreateNode(&node)

	// initialize the list of nodes for the new network
	var networkNodes swl.Nodes
	networkNodes = append(networkNodes, n)

	// Get the external ip address from the request
	netAddr := GetIPAddress(r)

	network := swl.Network{
		IP:    netAddr,
		Nodes: networkNodes,
		Stack: stack,
	}

	net := RepoCreateNetwork(&network)

	w.Header().Set("Content-Type", "application/json; charset=UTH-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(net); err != nil {
		panic(err)
	}
}

// NodeIndex sends a response with all current nodes in json format
func NodeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		panic(err)
	}
}

// NodeCreate creates a new node
// @param r: should contain 'name' which is the name of this new node and 'ip' which is the internal ip of this new node
// @param w: responds with 201 if succsesful
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

	if !swl.ValidateAndUnmarshalJSON(body, &node, w) {
		return
	}

	if !swl.ValidateParam("name", node.Name, w) {
		return
	}

	if !swl.ValidateParamRegex("ip", node.IP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w) {
		return
	}

	netAddr := GetIPAddress(r)

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

// NodeEnable sets the enable field for the specified node to enabled
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
		paramError := swl.ParamError{
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

// NodeGet returns info about a node given its id
// @param w: responds with the specified node if found, otherwise returns a 406
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
		paramError := swl.ParamError{
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

// NodeDelete removes the specified node from its network and from the list of all nodes
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

// NodeIncrementClients increases the Client field by 1
func NodeIncrementClients(w http.ResponseWriter, r *http.Request) {
	NodeUpdateClients(w, r, true)
}

// NodeDecrementClients decreases the Client field by 1
func NodeDecrementClients(w http.ResponseWriter, r *http.Request) {
	NodeUpdateClients(w, r, false)
}

// NodeUpdateClients either increments or decrements the number of clients for the node specified in r
func NodeUpdateClients(w http.ResponseWriter, r *http.Request, increment bool) {
	vars := mux.Vars(r)
	nodeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	if err := RepoUpdateNodeClients(nodeID, increment); err != nil {
		paramError := swl.ParamError{
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

// NetworkCurrent encodes the network object into JSON and sends it in the response
func NetworkCurrent(w http.ResponseWriter, r *http.Request) {
	netAddr := GetIPAddress(r)
	net := RepoFindNetworkByIP(netAddr)
	if net == nil {
		paramError := swl.ParamError{
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

// NetworkGetNodeDownload returns a node with the least amount of clients currently downloading from them
func NetworkGetNodeDownload(w http.ResponseWriter, r *http.Request) {
	netAddr := GetIPAddress(r)
	node := RepoFindBestNodeInNetworkByIP(netAddr)
	if node == nil {
		paramError := swl.ParamError{
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

// SoftwareGet gets best available node
func SoftwareGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	swl.DownloadLog.Info("Software Request " + vars["software_id"] + " Version " + vars["version_id"])
	swl.ConsoleLog.Info("Software Request " + vars["software_id"] + " Version " + vars["version_id"])

	NetworkGetNodeDownload(w, r)
}

// PackageGet gets best available node
func PackageGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	swl.DownloadLog.Info("Package Request " + vars["package_id"] + " Version " + vars["version_id"])
	swl.ConsoleLog.Info("Package Request " + vars["package_id"] + " Version " + vars["version_id"])
	NetworkGetNodeDownload(w, r)
}
