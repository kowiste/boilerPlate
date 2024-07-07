package repository

import (
	"boiler/src/model/user"
	"context"
)

type IUserRepository interface {
	CreateUser(c context.Context, user *user.User) (string, error)
	Users(c context.Context) (user.Users, error)
	UserByID(c context.Context, id string) (*user.User, error)
	UpdateUser(c context.Context, user *user.User) error
	DeleteUser(c context.Context, id string) error
}
