package seat

import (
	"MovieMate/db"
	"MovieMate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSeat(c *gin.Context) {
	var seat models.Seats

	if err := c.ShouldBindJSON(&seat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&seat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "seat added successfully"})

}
