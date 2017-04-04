package main

import (
	"net/http"

	"github.com/gorilla/mux"
	swl "github.com/stohio/software-lab/lib"
)

//NewRouter creates a new Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = swl.RouteLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
