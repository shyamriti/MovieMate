package routes

import (
	"MovieMate/controllers"
	"MovieMate/controllers/movie"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {
	r.POST("/api/login", controllers.Login)
}

func Signup(r *gin.Engine) {
	r.POST("/api/signup", controllers.Signup)
}

func AddMovie(r *gin.Engine) {
	r.POST("api/test", movie.AddMovie)
}