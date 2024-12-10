package validator

import (
	"ddd/shared/errors"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		
		// Register function to get json tag name
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	})
	return validate
}

// ValidateStruct validates a struct and returns AppError
func ValidateStruct(s interface{}) error {
	err := GetValidator().Struct(s)
	if err == nil {
		return nil
	}

	// Convert validator errors to AppError
	validationErrors := err.(validator.ValidationErrors)
	errorMessages := make([]string, 0, len(validationErrors))
	
	for _, e := range validationErrors {
		errorMessages = append(errorMessages, formatError(e))
	}
	
	return errors.NewValidation(
		strings.Join(errorMessages, "; "),
		err,
	)
}

// formatError formats a validation error into a readable message
func formatError(err validator.FieldError) string {
	field := err.Field()
	switch err.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "min":
		return field + " must be at least " + err.Param()
	case "max":
		return field + " must be at most " + err.Param()
	default:
		return field + " failed on " + err.Tag() + " validation"
	}
}