package main

import "time"

type Node struct {
	Id	int		`json:"id"`
	Name	*string		`json:"name"`
	LocalIP	*string		`json:"local_ip"`
	Network	string		`json:"network"`
	Enabled bool		`json:"enabled"`
	Added	time.Time	`json:"added"`
}

type Nodes []Node
