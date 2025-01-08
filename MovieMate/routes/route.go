package routes

import (
	"MovieMate/controllers"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {
	r.POST("/api/login", controllers.Login)
}

func Signup(r *gin.Engine) {
	r.POST("/api/signup", controllers.Signup)
}