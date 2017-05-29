package softwarelab

// Stack is a group of software and packages a user can download.
type Stack struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Packages  Packages  `json:"packages"`
	Softwares Softwares `json:"softwares"`
}

// Stacks records all the available Stacks a user can choose from.
type Stacks []*Stack

// Software is a particular application a user can download, either
// individually or in a stack. Also there may be different versions a
// user can download.
type Software struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Publisher string   `json:"publisher"`
	Versions  Versions `json:"versions"`
}

// Softwares keeps track of all a list of Software available in a Stack.
type Softwares []*Software

// Version is a particular form of some application. Dependent on the
// actual version number, the OS this specific software is for, as well
// as Architecture. Contains the url to download this software.
type Version struct {
	ID           int    `json:"id"`
	Version      string `json:"version"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	Extension    string `json:"extension"`
	URL          string `json:"url"`
	Checksum     string `json:"checksum"`
}

// Versions contains a list of Version structs for a particular Software
type Versions []*Version

// Package is a group of similar Software.
type Package struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Softwares   Softwares `json:"software"`
}

// Packages is a list of Package structs for a particular Stack.
type Packages []*Package
