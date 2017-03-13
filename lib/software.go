package softwarelab

type Stack struct {
	Id		int		`json:"id"`
	Name		string		`json:"name"`
	Packages	Packages	`json:"packages"`
	Softwares	Softwares	`json:"softwares"`

}

type Stacks []*Stack

type Software struct {
	Id		int		`json:"id"`
	Name		string		`json:"name"`
	Publisher	string		`json:"publisher"`
	Versions	Versions	`json:"versions"`
}

type Softwares []*Software

type Version struct {
	Id		int	`json:"id"`
	Version		string	`json:"version"`
	OS		string	`json:"os"`
	Architecture	string	`json:"architecture"`
	Extension	string	`json:"extension"`
	URL		string	`json:"url"`
}

type Versions []*Version

type Package struct {
	Id		int		`json:"id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Softwares	Softwares	`json:"software"`
}

type Packages []*Package
