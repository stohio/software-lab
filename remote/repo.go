package main

import "fmt"

var currentId int

var nodes Nodes

func init() {
	RepoCreateNode(Node {Name: "Node 1", LocalIP: "10.0.0.20", Network: "100.10.20.30"})
	RepoCreateNode(Node {Name: "Noded 2", LocalIP: "10.0.0.35", Network: "100.10.20.30"})

}

func RepoFindNode(id int) Node {
	for _, n := range nodes {
		if n.Id == id {
			return n
		}
	}
	//Otherwise, Return Empty
	return Node{}
}

func RepoCreateNode(n Node) Node {
	currentId += 1
	n.Id = currentId
	nodes = append(nodes, n)
	return n
}

func RepoDestroyNode(id int) error {
	for i, n := range nodes {
		if n.Id == id {
			nodes = append(nodes[:i], nodes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Node with id of %d to delete", id)
}

