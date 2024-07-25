package userservice

import (
	"boiler/src/model/user"
	"context"
)

func (serv UserService) Users(c context.Context, input *user.FindUsersInput) (users []user.User, err error) {
	return serv.db.Users(c, input)
}

func (serv UserService) UserByID(c context.Context) (users *user.User, err error) {
	return serv.db.UserByID(c, serv.user.ID)
}
