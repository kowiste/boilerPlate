package core

import (
	"boiler/pkg/validator"
	conf "boiler/src/config"
	"boiler/src/controller"
	"boiler/src/controller/grpc"
	"boiler/src/controller/kafka"
	"boiler/src/controller/rest"
	"boiler/src/repository"
	"boiler/src/repository/mysql"
	assetservice "boiler/src/service/asset"
	userservice "boiler/src/service/user"
	"fmt"

	"github.com/kowiste/config"
)

func Init() (err error) {

	cnf, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
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
	ctr := make([]controller.IController, 0)
	for i := range cnf.Controllers {
		switch cnf.Controllers[i] {
		case "rest":
			ctr = append(ctr, rest.New())
		case "grpc":
			ctr = append(ctr, grpc.New())
		case "kafka":
			ctr = append(ctr, kafka.New())
		}
		err = ctr[i].Init()
		if err != nil {
			return
		}
	}

	return
}
