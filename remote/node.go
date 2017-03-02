package main

import "time"

type Node struct {
	Name	string		`json:"name"`
	LocalIP	string		`json:"local_ip"`
	Network	string		`json:"network"`
	Added	time.Time	`json:"added"`
}

type Nodes []Node

