package main

import (
	api "deploy/api"
	"log"
	"net/http"
)

func main() {
	// Initialize routes
	api.SetupRoutes()

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
