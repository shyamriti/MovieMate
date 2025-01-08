package main

import (
	"MovieMate/config"
	"MovieMate/db"
	"MovieMate/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	config.LoadConfig()

	// Initialize the database connection
	db.Conn()

	if db.DB == nil {
		log.Fatal("Database connection is nil, exiting application")
	} else {
		log.Println("Database connected successfully")
	}

	// Initialize the Gin router
	route := gin.Default()

	// Add CORS middleware
	route.Use(cors.Default())

	// Define routes
	routes.Login(route)
	routes.Signup(route)
	routes.AddMovie(route)

	// Start the server on port 8080
	route.Run(":8080")
}
