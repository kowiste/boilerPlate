package db

import (
	"sync"

	"github.com/kowiste/boilerplate/pkg/errors"
)

type IDatabase interface {
	Init() error
	IUserDatabase
	IAssetDatabase
}

var (
	instance IDatabase
	once     sync.Once
)

// New returns the singleton instance of the repository
func New(injector IDatabase) IDatabase {
	once.Do(func() {
		instance = injector
	})
	return instance
}

// Get returns the singleton instance
func Get() (IDatabase, error) {
	if instance == nil {
		return nil, errors.New("repository not set", errors.EErrorServerInternal)
	}
	return instance, nil
}
