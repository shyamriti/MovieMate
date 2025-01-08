package main

import (
	"MovieMate/config"
	"MovieMate/db"
	"MovieMate/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db.Conn()

	if db.DB == nil {
		log.Fatal("Database connection is nil, exiting application")
	} else {
		log.Println("Database connected successfully")
	}

	route := gin.Default()
	routes.Login(route)
	routes.Signup(route)
	route.Run(":8080")
}
