package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is up")
}

func NodeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		panic(err)
	}
}
