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
		"VersionGet",
		"GET",
                "/download/software/{software_id:[0-9]+}/versions/{version_id:[0-9]+",
		SoftwareGet,
	},
	//Route {
	//	"GetPackage",
	//	"GET",
	//	"/",
	//	GetPackage,
	//},
}

