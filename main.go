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
	config.LoadConfig()

	db.Conn()

	if db.DB == nil {
		log.Fatal("Database connection is nil, exiting application")
	} else {
		log.Println("Database connected successfully")
	}

	route := gin.Default()
	route.Use(cors.Default())

	routes.Login(route)
	routes.Signup(route)
	routes.AddMovie(route)
	routes.GetMovies(route)
	routes.GetMovieByYear(route)
	routes.GetMovieByGenre(route)
	routes.AddSeat(route)
	routes.GetSeat(route)
	routes.UpdateSeat(route)
	routes.DeleteSeat(route)

	route.Run(":8080")
}
