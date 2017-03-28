package softwarelab

import "time"

// Node represents a local server on a particular network. Name, IP both refer to the cooresponding
// network this node is a part of
// TODO should IP refer to the local server machine, or the network the node is a part of?
type Node struct {
	Id      int       `json:"id"`
	Name    *string   `json:"name"`
	IP      *string   `json:"ip"`
	Enabled bool      `json:"enabled"`
	Clients int       `json:"clients"`
	Added   time.Time `json:"added"`
}

// Nodes is an array of Node structures
type Nodes []*Node
