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
		"NodeCreate",
		"POST",
		"/nodes",
		NodeCreate,
	},
}
