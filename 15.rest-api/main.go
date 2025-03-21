package main

import (
	"log"
	"net/http"
	"os"

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
	server.GET("/events", getEvents)
	server.Run(":" + port) // run the server on localhost:9999
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
