package main

import (
	"log"
	"net/http"
	"os"

	"chinsiang.com/rest-api/models"
	"chinsiang.com/rest-api/routes"
	"chinsiang.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	// please note that the .env file is not needed for the server to run
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set Gin mode from environment variable
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode // Default to DebugMode
	}
	gin.SetMode(mode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999" // Default port
	}

	server := gin.Default()
	server.SetTrustedProxies(nil)
	// server.SetTrustedProxies([]string{"192.168.1.0/24", "10.0.0.0/8"})
	server.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Register all routes
	routes.RegisterRoutes(server)

	// server.GET("/events", getEvents)
	// server.POST("/events", createEvent)
	server.Run(":" + port) // run the server on localhost:9999
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		validationErrors := utils.HandleValidationError(err, event.ValidationMessages())
		if len(validationErrors) > 0 {
			context.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		return
	}
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
