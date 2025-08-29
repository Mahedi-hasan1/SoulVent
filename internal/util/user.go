package util

import (
	"github.com/go-playground/validator/v10"
)

func GetCustomValidationMessage(err error) string {
	
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := map[string]map[string]string{
		"Password": {
			"required": "password is required",
			"min":      "password is less than 8 characters",
		},
		"Username": {
			"required": "username is required",
			"min":      "username must be at least 3 characters",
			"max":      "username cannot exceed 50 characters",
		},
		"Email": {
			"required": "email is required",
			"email":    "invalid email format",
		},
		"Gender": {
			"required": "gender is required",
			"oneof":    "gender must be male, female, or other",
		},
	}   
	for _, fieldError := range validationErrors {
		field := fieldError.Field()
		tag := fieldError.Tag()

		if fieldMessages, exists := errorMessages[field]; exists {
			if message, exists := fieldMessages[tag]; exists {
				return message
			}
		}
	}

	return err.Error() // fallback to default message
}
