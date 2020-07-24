package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PageNotFound handle 404
func PageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status": "Failed", "message": "WE ARE SORRY, BUT THE PAGE YOU REQUESTED WAS NOT FOUND"})
}

func WordCloud(c *gin.Context) {

	c.HTML(http.StatusOK, "cloud.html", gin.H{
		"payload": Decoded(),
		"top":     TopTen(),
		"cloud":   CloudChart(),
	})
}

func Radar(c *gin.Context) {

	// c.Header().Set("Content-Type", "text/html")
	c.HTML(http.StatusOK, "radar.html", gin.H{
		"payload": Decoded(),
		"top":     TopTen(),
		"cloud":   CloudChart(),
	})
}
