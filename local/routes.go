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
}
