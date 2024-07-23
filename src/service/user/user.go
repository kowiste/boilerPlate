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

func (serv UserService) Users(c context.Context, input *user.FindUsersInput) (users []user.User, err error) {
	return serv.db.Users(c, input)
}

func (serv UserService) UserByID(c context.Context) (users *user.User, err error) {
	return serv.db.UserByID(c, serv.user.ID)
}

func (serv UserService) Update(c context.Context) (err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.UpdateUser(c, serv.user)
}

// Delete removes a user from the database by ID.
// Parameters:
// - c: The context for the operation.
// - id: The user ID to delete.
// Returns:
// - err: An error if the deletion fails.
func (serv UserService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteUser(c, id)
}
