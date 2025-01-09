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
	r.POST("api/movie/add", movie.AddMovie)
}

func GetMovies(r *gin.Engine){
	r.GET("api/movies", movie.GetMovies)
}

func GetMovieByYear(r *gin.Engine){
	r.GET("api/movie", movie.GetMovieByYear)
}

func GetMovieByGenre(r *gin.Engine){
	r.GET("api/movie", movie.GetMovieByYear)
}