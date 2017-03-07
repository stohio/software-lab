package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"strconv"

	swl "github.com/stohio/software-lab/lib"
	"github.com/gorilla/mux"

)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up")
}

func NetworkIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(networks); err != nil {
		panic(err)
	}
}

func NetworkCreate(w http.ResponseWriter, r *http.Request) {
	var netCreate swl.NetworkCreate

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if !ValidateJson(body, &netCreate, w) {
		return
	}

	if !ValidateParam("name", netCreate.Name, w) {
		return
	}

	if !ValidateParamRegex("ip", netCreate.IP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w){
		return
	}


	stack := RepoFindStack(*netCreate.Stack)

	if stack == nil {
		response := ParamError {
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

	node := swl.Node {
		Name:	netCreate.Name,
		IP:	netCreate.IP,
	}
	n := RepoCreateNode(&node)

	var nodes swl.Nodes
	nodes = append(nodes, n)
	netAddr := GetIPAddress(r)

	network := swl.Network {
		IP:	netAddr,
		Nodes:	nodes,
		Stack:	stack,
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

	if !ValidateJson(body, &node, w) {
		return
	}

	if !ValidateParam("name", node.Name, w) {
		return
	}

	if !ValidateParamRegex("ip", node.IP, "\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b", w){
		return
	}

	netAddr  := GetIPAddress(r)

	if network := RepoFindNetworkByIP(netAddr); network != nil {
		n := RepoCreateNode(&node)
		network.Nodes = append(network.Nodes, n)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err:= json.NewEncoder(w).Encode(network); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(409)
		if err:= json.NewEncoder(w).Encode(stacks); err != nil {
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
		if err:= json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	} else {
		paramError := ParamError {
			Error: "Node with id not found",
			Param: "id",
			Value: vars["id"],
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(406)
		if err:= json.NewEncoder(w).Encode(paramError); err != nil {
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
		w.Header().Set("Content-Type", "application.json; charset=UTF-8")
		w.WriteHeader(200)
		if err:= json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	} else {
		paramError := ParamError {
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
		paramError := ParamError {
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
		paramError := ParamError {
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
        vars := mux.Vars(r)
	softwareID, err := strconv.Atoi(vars["software_id"])
	versionID, err := strconv.Atoi(vars["version_id"])
	node := RepoFindBestNodeInNetworkByIP(netAddr)
	if node == nil {
		paramError := ParamError {
			Error: "Could Not Find a Node",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(paramError); err != nil {
			panic(err)
		}
		return
	}

	client = &http.Client{}
        req, err := http.NewRequest("GET", node.IP + "/software/" + softwareID + "/verions/" + versionID)
        resp, err := client.Do(req)

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
