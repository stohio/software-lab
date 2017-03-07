package softwarelab

import "time"

type Node struct {
	Id	int		`json:"id"`
	Name	*string		`json:"name"`
	IP	*string		`json:"ip"`
	Enabled bool		`json:"enabled"`
	Clients	int		`json:"clients"`
	Added	time.Time	`json:"added"`
}

type Nodes []*Node
