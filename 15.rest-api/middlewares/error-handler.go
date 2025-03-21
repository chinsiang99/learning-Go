package middlewares

import (
	"encoding/json"
	"net/http"

	"chinsiang.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) == 0 {
			return
		}

		for _, err := range context.Errors {
			switch e := err.Err.(type) {
			case validator.ValidationErrors:
				// Safely retrieve custom validation messages
				customMessages, ok := context.Get("validationMessages")
				customFieldStructTagMapping, ok := context.Get("fieldStructTagMapping")
				if !ok {
					context.JSON(http.StatusInternalServerError, gin.H{"error": "Validation messages not provided"})
					return
				}

				// Ensure it's a map[string]string to prevent panic
				messageMap, ok := customMessages.(map[string]string)
				if !ok {
					context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation message format"})
					return
				}

				fieldStructTagMapping, ok := customFieldStructTagMapping.(map[string]string)
				if !ok {
					context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid field struct tag mapping format"})
					return
				}

				errors := utils.HandleValidationError(e, messageMap, fieldStructTagMapping)
				context.JSON(http.StatusBadRequest, gin.H{"errors": errors})
				return

			case *json.UnmarshalTypeError:
				context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data type provided"})
				return

			default:
				context.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
				return
			}
		}
	}
}
