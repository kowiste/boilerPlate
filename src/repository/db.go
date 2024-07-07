package repository

import (
	"boiler/pkg/errors"
	"sync"
)

type IRepository interface {
	Init() error
	IUserRepository
	IAssetRepository
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
		return nil, errors.New("repository not set", errors.EErrorServerInternal)
	}
	return instance, nil
}
