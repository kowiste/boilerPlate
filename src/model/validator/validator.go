package validator

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

//move to common package

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
		return nil, fmt.Errorf("repository not set")
	}
	return instance, nil
}
