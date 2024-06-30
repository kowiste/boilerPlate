package repository

import (
	"boiler/src/model/asset"
	"boiler/src/model/user"
	"context"
	"fmt"
	"sync"
)

type IRepository interface {
	Init() error
	//User
	CreateUser(context.Context, *user.User) (string, error)
	GetUsers(context.Context) (user.Users, error)
	GetUserByID(c context.Context, id string) (*user.User, error)
	UpdateUser(context.Context, *user.User) error
	DeleteUser(context.Context, string) error
	//Asset
	CreateAsset(context.Context, *asset.Asset) (string, error)
	GetAssets(context.Context) ([]asset.Asset, error)
	GetAssetByID(c context.Context, id string) (*asset.Asset, error)
	UpdateAsset(context.Context, *asset.Asset) error
	DeleteAsset(context.Context, string) error
}

var (
	instance IRepository
	once     sync.Once
)

// New returns the singleton instance of the repository
func New(injector IRepository) IRepository {
	once.Do(func() {
		instance = injector
	})
	return instance
}

// Get returns the singleton instance
func Get() (IRepository, error) {
	if instance == nil {
		return nil, fmt.Errorf("repository not set")
	}
	return instance, nil
}
