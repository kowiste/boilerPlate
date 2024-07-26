package core

import (
	"fmt"

	"github.com/kowiste/boilerplate/pkg/common/controller"
	"github.com/kowiste/boilerplate/pkg/validator"
	conf "github.com/kowiste/boilerplate/src/config"
	ownControl "github.com/kowiste/boilerplate/src/controller"
	"github.com/kowiste/boilerplate/src/controller/grpc"
	"github.com/kowiste/boilerplate/src/controller/kafka"
	"github.com/kowiste/boilerplate/src/controller/rest"
	"github.com/kowiste/boilerplate/src/repository"
	"github.com/kowiste/boilerplate/src/repository/mysql"
	assetservice "github.com/kowiste/boilerplate/src/service/asset"
	userservice "github.com/kowiste/boilerplate/src/service/user"

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
	//Init services
	assetservice.New(database)
	userservice.New(database)
	
	//Init controllers
	ctr := make([]ownControl.IController, 0)
	for i := range cnf.Controllers {
		switch cnf.Controllers[i] {
		case controller.Rest:
			ctr = append(ctr, rest.New())
		case controller.GRPC:
			ctr = append(ctr, grpc.New())
		case controller.Nats:
			ctr = append(ctr, kafka.New())
		}
		err = ctr[i].Init()
		if err != nil {
			return
		}
	}

	return
}
