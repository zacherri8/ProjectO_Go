package main

import (
	"log"
	"os"

	"projecto-backend/config"
	"projecto-backend/database"
	"projecto-backend/middleware"
	"projecto-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load .env variables
	config.LoadEnv()

	// Connect to PostgreSQL
	database.Connect()

	// Create Gin engine
	r := gin.Default()

	// GLOBAL MIDDLEWARE
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.RateLimitMiddleware())

	// Register all routes
	routes.RegisterRoutes(r)

	// Get port from .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port:", port)

	// Start server
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
