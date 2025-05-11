package api

import (
	"fmt"
	"net/http"
)

// Handler function for the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the API! by Fikri ")
}

// SetupRoutes initializes the routes for the API
func SetupRoutes() {
	http.HandleFunc("/", RootHandler)
}
