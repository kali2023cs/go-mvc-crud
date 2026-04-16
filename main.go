package main

import (
	"log"
	"os"
	"go-mvc-crud/config"
	"go-mvc-crud/routes"
	"go-mvc-crud/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default port")
	}

	// Connect to database
	config.ConnectDatabase()

	// Initialize Custom Validators
	utils.InitCustomValidators()

	// Setup router
	r := routes.SetupRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
