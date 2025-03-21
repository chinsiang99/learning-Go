package controllers

import (
	"fmt"
	"net/http"

	"chinsiang.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
		"events":  events,
	})
}

func CreateEvent(context *gin.Context) {
	var event models.Event

	// if err := context.ShouldBindJSON(&event); err != nil {
	// 	context.Error(err) // Pass error to middleware
	// 	return
	// }
	if err := context.ShouldBindJSON(&event); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Set custom error messages only for validation errors
			fmt.Println(event.ValidationMessages(), "this is validation message")
			context.Set("validationMessages", event.ValidationMessages())
			context.Set("fieldStructTagMapping", event.FieldStructTagMapping())
			context.Error(validationErrs) // Pass the error to middleware
			return
		}
		context.Error(err)
		return
	}

	// if err := event.Save(); err != nil {
	// 	context.Error(err)
	// 	return
	// }
	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}
