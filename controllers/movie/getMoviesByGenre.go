package movie

import (
	"MovieMate/db"
	"MovieMate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieByGenre(c *gin.Context) {
	var genre = c.Query("genre")
	var movie []models.Movie

	if genre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Genre"})
		return
	}

	if err := db.DB.Where("genre ILIKE  ?", genre).Find(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}
