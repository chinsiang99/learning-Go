package routes

import (
	"chinsiang.com/rest-api/controllers"
	"chinsiang.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
	// "github.com/yourusername/yourproject/controllers"
	// "github.com/yourusername/yourproject/middlewares"
)

func EventRoutes(router *gin.Engine) {
	events := router.Group("/events")
	// events.Use(middlewares.AuthMiddleware())
	events.Use(middlewares.ErrorHandler())
	{
		events.GET("/", controllers.GetEvents)
		events.POST("/", controllers.CreateEvent)
	}
}
