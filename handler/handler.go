// Package handler implements a series of functions to support registering pattern
// distinct handlers and delegate incoming request to them
//
package handler

import (
	"fmt"
	"log"
	"net/http"
)

var handlers = map[string]http.HandlerFunc{}

// Handle accepts incoming HTTP request and dispatch it to the corresponding handler function
func Handle(w http.ResponseWriter, r *http.Request) {
	if handler, matched := handlers[r.URL.Path]; matched {
		handler(w, r)
		return
	}

	w.WriteHeader(404)
	w.Write([]byte("NOT FOUND!"))
}

// Register registers the pattern and its corresponding handler function in the package's map storage
func Register(pattern string, handler http.HandlerFunc) {
	handlers[pattern] = handler
}

// Serve register its own Handle function to delegate all incoming HTTP requests,
// and then listen to the port 8000
func Serve() {
	http.HandleFunc("/", Handle)

	fmt.Println("Launching server listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
