package movie

import (
	"MovieMate/db"
	"MovieMate/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovieByYear(c *gin.Context) {
	var yearParam = c.Query("year")
	var movie []models.Movie

	year, err := strconv.Atoi(yearParam)
	if err != nil || year <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Year"})
		return
	}

	if err = db.DB.Where("year=?",year).Find(&movie).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}
