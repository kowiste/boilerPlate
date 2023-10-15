package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"test.com/config"
)

type validation struct {
	Validation *validator.Validate
}

var lock = &sync.Mutex{}
var singleInstance *validation

func New() error {
	if singleInstance == nil {
		singleInstance = new(validation)
		lock.Lock()
		defer lock.Unlock()
		singleInstance.Validation = validator.New()
		singleInstance.Validation.RegisterValidation("dbValidation", singleInstance.databaseValidation)
	}

	return nil
}

func GetInstance() *validation {
	return singleInstance
}

func (v validation) Validate(data any) error {
	return v.Validation.Struct(data)
}

func (v validation) databaseValidation(fl validator.FieldLevel) bool {
	if config.Get().DBType == "GORM" && !config.Get().DBMock {
		return config.Get().DBConnection != ""
	}
	return true
}
