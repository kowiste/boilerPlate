package user

import (
	"boiler/pkg/validator"
	"context"

	"github.com/google/uuid"
)

func (u *User) Validate(c context.Context) (err error) {
	u.ID = uuid.NewString()
	validate, err := validator.Get()
	if err!=nil{
		return
	}
	return validate.Struct(u)
}
