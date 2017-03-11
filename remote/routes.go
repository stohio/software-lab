package main

import (
	"net/http"
)

type Route struct {
	Name  string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"Test",
		"GET",
		"/",
		Test,
	},
	Route {
		"NodeIndex",
		"GET",
		"/nodes",
		NodeIndex,
	},
	Route {
		"NodeGet",
		"GET",
		"/nodes/{id:[0-9]+}",
		NodeGet,
	},
	Route {
		"NodeDelete",
		"DELETE",
		"/nodex/{id:[0-9]+}",
		NodeDelete,
	},
	Route {
		"NodeIncrementClients",
		"POST",
		"/nodes/{id:[0-9]+}/clients/increment",
		NodeIncrementClients,
	},
	Route {
		"NodeDecrementClients",
		"POST",
		"/nodes/{id:[0-9]+}/clients/decrement",
		NodeDecrementClients,
	},
	Route {
		"NodeCreate",
		"POST",
		"/nodes",
		NodeCreate,
	},
	Route {
		"NodeEnable",
		"POST",
		"/nodes/{id:[0-9]+}/enable",
		NodeEnable,
	},
	Route {
		"NetworkIndex",
		"GET",
		"/networks",
		NetworkIndex,
	},
	Route {
		"NetworkCurrent",
		"GET",
		"/networks/current",
		NetworkCurrent,
	},
	Route {
		"NetworkCreate",
		"POST",
		"/networks",
		NetworkCreate,
	},
	Route {
		"VersionGet",
		"GET",
		"/software/{software_id:[0-9]+}/versions/{version_id:[0-9]+}",
		SoftwareGet,
	},
	Route {
		"PackageGet",
		"GET",
		"/packages/{package_id:[0-9]+}/versions/{version_id:[0-9]+}",
		PackageGet,
	},
}

