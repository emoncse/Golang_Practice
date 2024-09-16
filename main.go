package main

import (
	"first_project/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting the server on :8080...")

	// Register routes
	router := routes.RegisterRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
