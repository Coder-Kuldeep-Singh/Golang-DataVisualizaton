package routes

import (
	"Dashboard/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("html/templates/*")
	router.Static("/home", "./html/css")

	router.GET("/", controllers.TableDataSend)
	// router.GET("/role/:role", controllers.GetHomeWithRole)
	// router.GET("/tnc", controllers.GetTNC)
	return router
}
