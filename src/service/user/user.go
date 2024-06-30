package userservice

import (
	"boiler/src/model/user"
	"boiler/src/repository"
	"context"
)

type UserService struct {
	user *user.User
	db   repository.IRepository
}

func New() (serv *UserService, err error) {
	database, err := repository.Get()
	if err != nil {
		return
	}
	return &UserService{
		user: new(user.User),
		db:   database,
	}, nil
}
func (serv *UserService) GetUser() *user.User {
	return serv.user
}

func (serv UserService) Create(c context.Context) (id string, err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.CreateUser(c, serv.user)
}

func (serv UserService) Get(c context.Context) (users []user.User, err error) {
	return serv.db.GetUsers(c)
}

func (serv UserService) GetByID(c context.Context) (users *user.User, err error) {
	return serv.db.GetUserByID(c, serv.user.ID)
}

func (serv UserService) Update(c context.Context) (err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.UpdateUser(c, serv.user)
}
func (serv UserService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteUser(c, id)
}
