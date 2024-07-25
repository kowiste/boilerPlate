package userservice

import (
	"context"
)

func (serv UserService) Create(c context.Context) (id string, err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.CreateUser(c, serv.user)
}
