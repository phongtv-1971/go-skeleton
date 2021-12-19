package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
		"success": false,
	}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"data": data,
		"success": true,
	}
}

func ValidationErrorResponse(err error) gin.H  {
	var validations []Validation
	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field(), value.Tag()
		validation := Validation{Field: field, Message: generateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}

	return gin.H{
		"error": validations,
		"success": false,
	}
}

func generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field %s is %s.", field, rule)
	case "email":
		return fmt.Sprintf("Field %s is invalid format.", field)
	default:
		return fmt.Sprintf("Field %s is not valid.", field)
	}
}
