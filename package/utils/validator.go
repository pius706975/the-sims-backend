package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errorsMap := make(map[string]string)
		for _, fe := range ve {
			field := toSnakeCase(fe.Field())
			switch fe.Tag() {
			case "required":
				errorsMap[field] = field + " is required"
			default:
				errorsMap[field] = field + " is invalid"
			}
		}
		return errorsMap
	}

	return map[string]string{
		"error": err.Error(),
	}
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
