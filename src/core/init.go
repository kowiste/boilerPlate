package core

import (
	"boiler/src/controller"
	"boiler/src/controller/rest"
	"boiler/src/model/validator"
	"boiler/src/repository"
	"boiler/src/repository/mysql"
	"fmt"
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
	var ctr controller.IController = rest.New()
	err = ctr.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
