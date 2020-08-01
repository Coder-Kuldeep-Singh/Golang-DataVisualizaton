package controllers

import (
	"fmt"
	"net/http"

	"Dashboard/models"

	"github.com/gin-gonic/gin"
)

func ReadUserIP(c *gin.Context) string {
	IPAddress := c.Request.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = c.Request.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = c.Request.RemoteAddr
	}
	return IPAddress
}

//PageNotFound handle 404
func PageNotFound(c *gin.Context) {
	// ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	// fmt.Println("Proxy IP address", ip)
	// RealIP := c.Request.Header.Get("X-FORWARDED-FOR")
	// fmt.Println("Clients Real IP Addres :", RealIP)
	ip := ReadUserIP(c)
	fmt.Println("User IP :", ip)
	c.JSON(http.StatusNotFound, gin.H{"status": "Failed", "message": "WE ARE SORRY, BUT THE PAGE YOU REQUESTED WAS NOT FOUND"})
	models.UserIPAddress(ip)

}

func WordCloud(c *gin.Context) {
	// ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	// fmt.Println("Proxy IP address", ip)
	// RealIP := c.Request.Header.Get("X-FORWARDED-FOR")
	// fmt.Println("Clients Real IP Addres :", RealIP)
	ip := ReadUserIP(c)
	fmt.Println("User IP :", ip)
	c.HTML(http.StatusOK, "cloud.html", gin.H{
		"payload": Decoded(),
		"top":     TopTen(),
		"cloud":   CloudChart(),
	})
	models.UserIPAddress(ip)

}

func Radar(c *gin.Context) {
	// ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	// fmt.Println("Proxy IP address", ip)
	// RealIP := c.Request.Header.Get("X-FORWARDED-FOR")
	// fmt.Println("Clients Real IP Addres :", RealIP)
	ip := ReadUserIP(c)
	fmt.Println("User IP :", ip)

	// c.Request.Header().Set("Content-Type", "text/html")
	c.HTML(http.StatusOK, "radar.html", gin.H{
		"payload": Decoded(),
		"top":     TopTen(),
		"cloud":   CloudChart(),
	})
	models.UserIPAddress(ip)
}
