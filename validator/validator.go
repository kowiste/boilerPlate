package validator

import "github.com/go-playground/validator/v10"

func New(data any) error {
	validate := validator.New()
	return validate.Struct(data)
}
