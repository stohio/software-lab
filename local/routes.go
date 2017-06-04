package main

import (
	swl "github.com/stohio/software-lab/lib"
)

var routes = swl.Routes{
	swl.Route{
		Name:        "Test",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Test,
	},
	swl.Route{
		Name:        "VersionGet",
		Method:      "GET",
		Pattern:     "/download/software/{software_id:[0-9]+}/versions/{version_id:[0-9]+}",
		HandlerFunc: SoftwareGet,
	},
	swl.Route{
		Name:        "PackageGet",
		Method:      "GET",
		Pattern:     "/download/package/{package_id:[0-9]+}/versions/{version_id:[0-9]+}",
		HandlerFunc: PackageGet,
	},
	swl.Route{
		Name:        "NodeGet",
		Method:      "GET",
		Pattern:     "/node",
		HandlerFunc: NodeGet,
	},
	swl.Route{
		Name:        "NodeUpdate",
		Method:      "UPDATE",
		Pattern:     "/node",
		HandlerFunc: NodeUpdate,
	},
	swl.Route{
		Name:        "Stacks",
		Method:      "GET",
		Pattern:     "/stacks",
		HandlerFunc: StacksGet,
	},
	swl.Route{
		Name:        "NetworkGet",
		Method:      "GET",
		Pattern:     "/network",
		HandlerFunc: NetworkGet,
	},
	swl.Route{
		Name:        "NetworkPost",
		Method:      "POST",
		Pattern:     "/network",
		HandlerFunc: NetworkPost,
	},
	swl.Route{
		Name:        "NodeInit",
		Method:      "POST",
		Pattern:     "/node/init",
		HandlerFunc: NodeInit,
	},
	swl.Route{
		Name:        "NodeInitStatus",
		Method:      "GET",
		Pattern:     "/node/init",
		HandlerFunc: NodeInitStatus,
	},
}
