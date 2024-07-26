package userservice

import (
	"sync"

	"github.com/kowiste/boilerplate/pkg/errors"
	"github.com/kowiste/boilerplate/src/model/user"
	"github.com/kowiste/boilerplate/src/repository"
)

type UserService struct {
	user *user.User
	db   repository.IRepository
}

var (
	instance *UserService
	once     sync.Once
)

func New(db repository.IRepository) (serv *UserService) {

	once.Do(func() {

		instance = &UserService{
			user: new(user.User),
			db:   db,
		}
	})
	return instance
}

func Get() (*UserService, error) {
	if instance == nil {
		return nil, errors.New("AssetService not set", errors.EErrorServerInternal)
	}
	return instance, nil
}
func (serv *UserService) GetUser() *user.User {
	return serv.user
}
