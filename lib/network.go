package softwarelab

// Network represents a group of local servers, and the stack of applications that they offer
// TODO Id should be renamed to ID
type Network struct {
	Id    int    `json:"id"`
	IP    string `json:"ip"`
	Nodes Nodes  `json:"nodes"`
	Stack *Stack `json:"stack"`
}

// Networks is a list of all the current Networks set up
type Networks []*Network

// NetworkCreate is a model of the JSON in the body of a request to create a new network
// IP represents the local IP of the fist node in the network
// Name represents the name of the first node in the network
// Stack is an id for the stack that will be used in the network
// TODO Stack should be renamed to StackId
type NetworkCreate struct {
	IP    *string `json:"ip"`
	Name  *string `json:"name"`
	Stack *int    `json:"stack"`
}
