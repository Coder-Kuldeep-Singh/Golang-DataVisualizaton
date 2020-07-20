package routes

import (
	"Dashboard/controllers"
	"Dashboard/services"

	"github.com/gin-gonic/gin"
)

//SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("html/templates/*")
	router.Static("/home", "./html/css")

	router.GET("/", controllers.TableDataSend)
	api := router.Group("/api")
	{
		api.GET("/json/data", services.GetAll)
	}
	// router.GET("/role/:role", controllers.GetHomeWithRole)
	router.NoRoute(controllers.PageNotFound)
	return router
}
