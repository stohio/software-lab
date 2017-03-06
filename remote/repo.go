package main

import (
	"fmt"
	"time"

	swl "github.com/stohio/software-lab/lib"
)

var currentNodeId int
var currentNetworkId int
var currentStackId int

var nodes swl.Nodes
var networks swl.Networks

var stacks swl.Stacks

func init() {

	version1 := swl.Version {
		Version:	"1.0",
		OS:		"Windows",
		Architecture:	"64",
		URL:		"http://www.textfiles.com/humor/failure.txt",
	}

	version2 := swl.Version {
		Version:	"1.0",
		OS:		"Mac",
		Architecture:	"64",
		URL:		"http://www.textfiles.com/humor/failure.txt",
	}

	var versions swl.Versions
	versions = append(versions, &version1)
	versions = append(versions, &version2)

	software1 := swl.Software {
		Id:		1,
		Name:		"My Software",
		Publisher:	"Stohio",
		Versions:	versions,
	}

	software2 := swl.Software {
		Id:		2,
		Name:		"Not My Software",
		Publisher:	"Stohio",
		Versions:	versions,
	}

	var softwares swl.Softwares
	softwares = append(softwares, &software1)
	softwares = append(softwares, &software2)

	package1 := swl.Package {
		Id:		1,
		Name:		"Pack One",
		Description:	"This is a Package",
		Softwares:	softwares,
	}

	package2 := swl.Package {
		Id:		2,
		Name:		"Two Pack",
		Description:	"This is also a Package",
		Softwares:	softwares,
	}
	var packages swl.Packages
	packages = append(packages, &package1)
	packages = append(packages, &package2)


	stack := swl.Stack {
		Id:		1,
		Name:		"My First Stack",
		Packages:	packages,
		Softwares:	softwares,
	}
	stacks = append(stacks, &stack)

}

func RepoFindStack(id int) *swl.Stack {
	for _, s :=range stacks {
		if s.Id == id {
			return s
		}
	}
	return nil
}

func RepoCreateStack(s *swl.Stack) *swl.Stack {
	currentStackId += 1
	s.Id = currentStackId
	stacks = append(stacks, s)
	return s
}

func RepoDestroyStack(id int) error {
	for i, s := range stacks {
		if s.Id == id {
			stacks = append(stacks[:i], stacks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Stack with id of %d to delete", id)
}

func RepoFindNetworkByIP(ip string) (*swl.Network, error) {
	for _, net := range networks {
		if net.IP == ip {
			return net, nil
		}
	}
	return nil,fmt.Errorf("Unable to find Network with ip of %s", ip)
}

func RepoCreateNetwork(n *swl.Network) *swl.Network {
	currentNetworkId += 1
	n.Id = currentNetworkId
	networks = append(networks, n)
	return n
}

func RepoDestroyNetwork(id int) error {
	for i, n := range networks {
		if n.Id == id {
			networks = append(networks[:i], networks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Network with id of %d to delete", id)
}

func RepoFindNode(id int) *swl.Node {
	for _, n := range nodes {
		if n.Id == id {
			return n
		}
	}
	//Otherwise, Return Empty
	return nil
}

func RepoCreateNode(n *swl.Node) *swl.Node {
	currentNodeId += 1
	n.Id = currentNodeId
	n.Added = time.Now()
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

