package main

import (
	swl "github.com/stohio/software-lab/lib"
)

var network swl.Network
var node swl.Node

// RepoFindSoftwareVersion gets the software version specified, if it exists
// @param id: software id
// @param version: version id
// @return: the version with the software id / version id specified, nil  if the version can't be found
func RepoFindSoftwareVersion(id int, version int) *swl.Version {
	for _, s := range network.Stack.Softwares {
		if s.ID == id {
			for _, v := range s.Versions {
				if v.ID == version {
					return v
				}
			}
		}
	}
	return nil
}

// RepoFindPackage gets the package specified by ID, if it exists
// @param id: package id
// @return: the package with the package id, nil if the package can't be found
func RepoFindPackage(id int) *swl.Package {
	for _, p := range network.Stack.Packages {
		if p.ID == id {
			return p
		}
	}
	return nil
}
