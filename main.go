package main

import (
	"first_project/database"
	"first_project/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load environment variables from .env file
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("Starting the server on :8080...")

	database.InitializeDB()

	// Register routes
	router := routes.RegisterRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
