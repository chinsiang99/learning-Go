package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ErrorResponse defines the structure of validation error responses
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// HandleValidationError processes validation errors with custom messages
func HandleValidationError(err error, customMessages map[string]string) []ErrorResponse {
	var validationErrors []ErrorResponse

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			// Check if a custom message is provided
			fmt.Print(e.Field())
			message, found := customMessages[e.Field()]
			if !found {
				message = "Invalid input for " + e.Field()
			}

			validationErrors = append(validationErrors, ErrorResponse{
				Field:   e.Field(),
				Message: message,
			})
		}
	}
	return validationErrors
}
