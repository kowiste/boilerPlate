package service

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (s service) validate(err error) (bool, int, map[string]string) {
	errorMessages := make(map[string]string)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, errorMessage := range validationErrors {
				errorMessages[errorMessage.Namespace()] = errorMessage.Error()
			}
			return false, http.StatusUnprocessableEntity, errorMessages
		} else {
			println(err.Error())
			return false, http.StatusBadRequest, errorMessages
		}
	}
	return true, http.StatusOK, errorMessages
}
