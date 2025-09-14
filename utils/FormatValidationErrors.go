package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationErrors converts validator.ValidationErrors to user-friendly messages
func FormatValidationErrors(errs validator.ValidationErrors) map[string][]string {
	messages := make(map[string][]string)
	for _, e := range errs {
		var msg string
		switch e.Tag() {
		case "required":
			msg = e.Field() + " is required"
		case "email":
			msg = "Email format invalid"
		case "unique_email":
			msg = "Email already exists"
		case "gte":
			msg = e.Field() + " must be at least " + e.Param() + " characters"
		case "lte":
			msg = e.Field() + " must be at most " + e.Param() + " characters"
		case "number":
			msg = e.Field() + " must be a number"
		default:
			msg = e.Field() + ": " + e.Error()
		}
		messages[strings.ToLower(e.Field())] = append(messages[strings.ToLower(e.Field())], msg)
	}
	return messages
}
