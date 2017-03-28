package softwarelab

// Network represents a group of local servers, and the stack of applications that they offer
type Network struct {
	Id    int    `json:"id"`
	IP    string `json:"ip"`
	Name  string `json:"name"`
	Nodes Nodes  `json:"nodes"`
	Stack *Stack `json:"stack"`
}

// Networks is a list of all the current Networks set up
type Networks []*Network

// TODO why does NetworkCreate have Name but Network doesn't?
// NetworkCreate is a potential Network structure
type NetworkCreate struct {
	IP    *string `json:"ip"`
	Name  *string `json:"name"`
	Stack *int    `json:"stack"`
}
