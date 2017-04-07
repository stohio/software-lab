package softwarelab

import "net/http"

// Route represents a link between a url path and a handler function found in handlers.go
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of Route structures
type Routes []Route
