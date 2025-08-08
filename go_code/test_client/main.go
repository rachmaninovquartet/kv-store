package main

import (
	"log"
	"os"

	"test_client/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Gin
	r := gin.Default()

	// Setup routes
	api.SetupRoutes(r)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	log.Printf("Starting test client on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start test client:", err)
	}
}
