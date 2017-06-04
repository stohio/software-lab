package softwarelab

// Network represents a group of local servers, and the stack of applications that they offer
type Network struct {
	ID    int    `json:"id"`
	IP    string `json:"ip"`
	Nodes Nodes  `json:"nodes"`
	Stack *Stack `json:"stack"`
}

// Networks is a list of all the current Networks set up
type Networks []*Network

// NetworkCreate is a model of the JSON in the body of a request to create a new network
// Stack is an id for the stack that will be used in the network
type NetworkCreate struct {
	StackID *int `json:"stack"`
}
