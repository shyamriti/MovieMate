package seat

import (
	"MovieMate/db"
	"MovieMate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSeat(c *gin.Context) {
	var seat models.Seats

	var seatClass = c.Query("class")

	if seatClass == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seat class"})
		return
	}

	if err := db.DB.Where("seat_class = ?", seatClass).Find(&seat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seat)
}
