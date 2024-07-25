package userservice

import (
	"context"
)

func (serv UserService) Update(c context.Context) (err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.UpdateUser(c, serv.user)
}
