package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TableDataSend(c *gin.Context) {
	// titles := AllTitles()
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"payload": Decoded(),
		"top":     TopTen(),
	})
}
