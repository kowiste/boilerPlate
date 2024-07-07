package validator

import (
	"boiler/pkg/errors"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	instance *validator.Validate
	once     sync.Once
)

func New() *validator.Validate {
	once.Do(func() {
		instance = validator.New(validator.WithRequiredStructEnabled())
	})
	return instance
}
func Get() (*validator.Validate, error) {
	if instance == nil {
		return nil, errors.New("validator not set", errors.EErrorServerInternal)
	}
	return instance, nil
}
