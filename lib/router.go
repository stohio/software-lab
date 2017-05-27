package softwarelab

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router that maps all routes in routes to their respective handler function
func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = RouteLogger(handler, route.Name)

		router.PathPrefix("/softwarelab").Handler(http.StripPrefix("/softwarelab/", http.FileServer(http.Dir("../site/"))))

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
