package seat

import (
	"MovieMate/db"
	"MovieMate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSeat(c *gin.Context) {
	var seat []models.Seats

	if err := db.DB.Find(&seat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seat)
}
