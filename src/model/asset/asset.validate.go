package asset

import (
	"boiler/pkg/validator"
	"context"

	"github.com/google/uuid"
)

func (a *Asset) Validate(c context.Context) (err error) {
	a.ID = uuid.NewString()
	validate, err := validator.Get()
	if err != nil {
		return
	}
	return validate.Struct(a)
}
