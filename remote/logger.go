package main

import (
	"log"
	"net/http"
	"time"
)

// Logger prints debug information for a route handler function
// @param inner: the handler object to be logged
// @param hame: the name of the route this handler is associated with
// @return: the new handler object with the modified handler function to include logging
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
