package db

import (
	"boiler/src/model/user"
	"fmt"
)

type IDatabase interface {
	Init() error
	CreateUser(*user.User) error
	GetUsers() (*user.Users, error)
}

//Move this to a share place

var instance IDatabase

func New(injector IDatabase) IDatabase {
	instance = injector
	return instance
}

func Get() (db IDatabase, err error) {
	if instance == nil {
		return instance, fmt.Errorf("database not set")
	}
	return instance, nil
}
