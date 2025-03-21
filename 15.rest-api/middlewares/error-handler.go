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

		// Retrieve custom messages if available
		customMessages, _ := context.Get("validationMessages")

		for _, err := range context.Errors {
			switch e := err.Err.(type) {
			case validator.ValidationErrors:
				messages := map[string]string{}
				if customMessages != nil {
					messages = customMessages.(map[string]string)
				}
				errors := utils.HandleValidationError(e, messages)
				context.JSON(http.StatusBadRequest, gin.H{"errors": errors})
				return

			case *json.UnmarshalTypeError:
				context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data type provided"})
				return

			default:
				context.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
				return
			}
		}
	}
}
