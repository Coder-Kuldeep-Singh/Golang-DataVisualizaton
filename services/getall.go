package services

import (
	"Dashboard/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	display := controllers.Decoded()
	// fmt.Println(len(display))

	if len(display) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "Failed", "error": "No records found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "Visual": display})
	}
}
