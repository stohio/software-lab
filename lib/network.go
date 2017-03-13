package softwarelab

type Network struct {
	Id	int	`json:"id"`
	IP	string	`json:"ip"`
	Nodes	Nodes	`json:"nodes"`
	Stack	*Stack	`json:"stack"`
}

type Networks []*Network


type NetworkCreate struct {
	IP	*string	`json:"ip"`
	Name	*string	`json:"name"`
	Stack	*int	`json:"stack"`
}
