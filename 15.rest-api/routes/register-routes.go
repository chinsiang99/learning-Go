package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Public routes
	// routes.AuthRoutes(router)
	// routes.UserRoutes(router)

	// Protected routes
	EventRoutes(router)
}
