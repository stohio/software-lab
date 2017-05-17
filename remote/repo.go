package main

import (
	"encoding/json"
	"fmt"
	"time"

	swl "github.com/stohio/software-lab/lib"
)

var currentNodeID int
var currentNetworkID int
var currentStackID int

var nodes swl.Nodes
var networks swl.Networks

var stacks swl.Stacks

func init() {

	var stack swl.Stack

	jsonStack := `{
  "id": 1,
  "name": "Hackathon",
  "softwares": [
    {
      "id": 1,
      "name": "Node JS",
      "publisher": "Node",
      "versions": [
        {
          "id": 1,
          "version": "6.10.0",
          "os": "Windows",
          "architecture": "32",
          "extension": ".msi",
          "url": "https://nodejs.org/dist/v6.10.0/node-v6.10.0-x86.msi"
        },
        {
          "id": 2,
          "version": "6.10.0",
          "os": "Mac",
          "architecture": "64",
          "extension": ".pkg",
          "url": "https://nodejs.org/dist/v6.10.0/node-v6.10.0.pkg"
        }
      ]
    },
    {
      "id": 2,
      "name": "Python",
      "publisher": "Python",
      "versions": [
        {
          "id": 1,
          "version": "3.6.0",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://www.python.org/ftp/python/3.6.0/python-3.6.0.exe"
        },
        {
          "id": 2,
          "version": "3.6.0",
          "os": "Mac",
          "architecture": "64",
          "extension": ".pkg",
          "url": "https://www.python.org/ftp/python/3.6.1/python-3.6.1rc1-macosx10.6.pkg"
        }
      ]
    }
  ]
}`

	if err := json.Unmarshal([]byte(jsonStack), &stack); err != nil {
		panic(err)
	}

	var softs swl.Softwares
	softs = append(softs, stack.Softwares[0])
	softs = append(softs, stack.Softwares[1])
	pack := swl.Package{
		ID:          1,
		Name:        "My First Package",
		Description: "This is the first package",
		Softwares:   softs,
	}
	stack.Packages = append(stack.Packages, &pack)
	stacks = append(stacks, &stack)

}

// RepoFindStack returns the stack associated with an id
// @param id: the id of the stack to get
// @return: returns a pointer to the stack with the id id
func RepoFindStack(id int) *swl.Stack {
	for _, s := range stacks {
		if s.ID == id {
			return s
		}
	}
	return nil
}

// RepoCreateStack takes a stack structure and adds this stack to the list of stacks
// @param s: a pointer to the newly created stack
// @return: returns the stack with the its ID set
func RepoCreateStack(s *swl.Stack) *swl.Stack {
	currentStackID++
	s.ID = currentStackID
	stacks = append(stacks, s)
	return s
}

// RepoDestroyStack removes a stack with the specified id
// @param id: the id of the stack to destroy
// @return: returns nil if successful and returns and error if it can't find the stack
func RepoDestroyStack(id int) error {
	for i, s := range stacks {
		if s.ID == id {
			stacks = append(stacks[:i], stacks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Stack with id of %d to delete", id)
}

// RepoFindNetworkByIP gets the network with the specified ip
// @param ip: the ip of the network to get
// @return: the network with the ip specified, nil if the network can't be found
func RepoFindNetworkByIP(ip string) *swl.Network {
	for _, net := range networks {
		if net.IP == ip {
			return net
		}
	}
	return nil
}

// RepoFindBestNodeInNetworkByIP gets the node in the network specified by ip with the smallest number of clients
// @param ip: the ip of the network to search in
// @return: the node with the least amount of clients in the specified network
func RepoFindBestNodeInNetworkByIP(ip string) *swl.Node {
	net := RepoFindNetworkByIP(ip)
	if net == nil {
		fmt.Println("Could Not Find Network")
		return nil
	}
	var bestNode *swl.Node
	bestDownloads := -1
	for _, n := range net.Nodes {
		fmt.Printf("Node, Best: %d. %d\n", n.Clients, bestDownloads)
		if (n.Clients < bestDownloads || bestDownloads == -1) && (n.Enabled) {
			fmt.Println("Best Node Updated!")
			bestNode = n
			bestDownloads = n.Clients
			if bestDownloads == 0 {
				return bestNode
			}
		}
	}
	if bestDownloads == -1 {
		return nil
	}
	return bestNode

}

// RepoCreateNetwork takes in a network struct and adds it to the list of all networks
// @param n: a Network struct
// @return: the network that was just created
func RepoCreateNetwork(n *swl.Network) *swl.Network {
	currentNetworkID++
	n.ID = currentNetworkID
	fmt.Println("Added Network")
	networks = append(networks, n)
	return n
}

// RepoDestroyNetwork deletes a network witht eh given id
// @param id: the id of the network to destroy
// @return: nil if successful and an error if the network with the given id doesn't exist
func RepoDestroyNetwork(id int) error {
	for i, n := range networks {
		if n.ID == id {
			networks = append(networks[:i], networks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Network with id of %d to delete", id)
}

// RepoFindNode returns a node with the given id
// @param id: the id of node to get
// @return: the node with the given ip or nil
func RepoFindNode(id int) *swl.Node {
	for _, n := range nodes {
		if n.ID == id {
			return n
		}
	}
	//Otherwise, Return Empty
	return nil
}

// RepoCreateNode adds the given node to list of nodes and sets its id value
// @param n: the new node to add
// @return: the new node with updated id
func RepoCreateNode(n *swl.Node) *swl.Node {
	currentNodeID++
	n.ID = currentNodeID
	n.Added = time.Now()
	nodes = append(nodes, n)
	return n
}

// RepoEnableNode sets the node with given id to enabled
// @param id: the id of the node to update
// @return: returns nil if the node isn't found, otherwise returns the updated node
func RepoEnableNode(id int) *swl.Node {
	node := RepoFindNode(id)
	if node == nil {
		return nil
	}
	node.Enabled = true
	return node
}

// DeleteNode removes the node with the given id
// @param id: the id of node to delete
// @return: returns nil if successful, otherwise returns an Error
func DeleteNode(id int) error {
	for _, n := range networks {
		for j, nod := range n.Nodes {
			if nod.ID == id {
				n.Nodes = append(n.Nodes[:j], n.Nodes[j+1:]...)
				fmt.Println(len(n.Nodes))
				if len(n.Nodes) == 0 {
					RepoDestroyNetwork(n.ID)
				}
				break
			}
		}
	}
	for i, n := range nodes {
		if n.ID == id {
			nodes = append(nodes[:i], nodes[i+1:]...)
			return nil
		}

	}
	return fmt.Errorf("Unable to find Node with id of %d to delete", id)
}

// RepoUpdateNodeClients either increments or decrements the Client field of the node with the given id
// @param id: the id of the node to update
// @param increment: if true the client field is incremented by 1, if false its decremented by 1
// @return: nil on success, otherwise an Error
func RepoUpdateNodeClients(id int, increment bool) error {
	for _, n := range nodes {
		if n.ID == id {
			if increment {
				n.Clients++
			} else {
				n.Clients++
			}
			return nil
		}
	}
	return fmt.Errorf("Unable to find Node with id of %d to Update Clients", id)
}
