package core

import (
	"boiler/pkg/validator"
	"boiler/src/controller"
	"boiler/src/controller/rest"
	"boiler/src/repository"
	"boiler/src/repository/mysql"
	assetservice "boiler/src/service/asset"
	userservice "boiler/src/service/user"
)

func Init() (err error) {
	//Init Validator
	validator.New()
	//Init database
	repository.New(mysql.New())
	database, err := repository.Get()
	if err != nil {
		return
	}
	err = database.Init()
	if err != nil {
		return
	}
	assetservice.New(database)
	userservice.New(database)
	var ctr controller.IController = rest.New()
	err = ctr.Init()
	if err != nil {
		return
	}

	return
}
