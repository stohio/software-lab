package main

import swl "github.com/stohio/software-lab/lib"

var routes = swl.Routes{
	swl.Route{
		Name:        "Test",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Test,
	},
	swl.Route{
		Name:        "NodeIndex",
		Method:      "GET",
		Pattern:     "/nodes",
		HandlerFunc: NodeIndex,
	},
	swl.Route{
		Name:        "NodeGet",
		Method:      "GET",
		Pattern:     "/nodes/{id:[0-9]+}",
		HandlerFunc: NodeGet,
	},
	swl.Route{
		Name:        "NodeDelete",
		Method:      "DELETE",
		Pattern:     "/nodes/{id:[0-9]+}",
		HandlerFunc: NodeDelete,
	},
	swl.Route{
		Name:        "NodeIncrementClients",
		Method:      "POST",
		Pattern:     "/nodes/{id:[0-9]+}/clients/increment",
		HandlerFunc: NodeIncrementClients,
	},
	swl.Route{
		Name:        "NodeDecrementClients",
		Method:      "POST",
		Pattern:     "/nodes/{id:[0-9]+}/clients/decrement",
		HandlerFunc: NodeDecrementClients,
	},
	swl.Route{
		Name:        "NodeCreate",
		Method:      "POST",
		Pattern:     "/nodes",
		HandlerFunc: NodeCreate,
	},
	swl.Route{
		Name:        "NodeEnable",
		Method:      "POST",
		Pattern:     "/nodes/{id:[0-9]+}/enable",
		HandlerFunc: NodeEnable,
	},
	swl.Route{
		Name:        "NetworkIndex",
		Method:      "GET",
		Pattern:     "/networks",
		HandlerFunc: NetworkIndex,
	},
	swl.Route{
		Name:        "NetworkCurrent",
		Method:      "GET",
		Pattern:     "/networks/current",
		HandlerFunc: NetworkCurrent,
	},
	swl.Route{
		Name:        "NetworkCreate",
		Method:      "POST",
		Pattern:     "/networks",
		HandlerFunc: NetworkCreate,
	},
	swl.Route{
		Name:        "VersionGet",
		Method:      "GET",
		Pattern:     "/software/{software_id:[0-9]+}/versions/{version_id:[0-9]+}",
		HandlerFunc: SoftwareGet,
	},
	swl.Route{
		Name:        "PackageGet",
		Method:      "GET",
		Pattern:     "/packages/{package_id:[0-9]+}/versions/{version_id:[0-9]+}",
		HandlerFunc: PackageGet,
	},
}
