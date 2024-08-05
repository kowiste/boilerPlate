package core

import (
	"fmt"

	"github.com/kowiste/boilerplate/src/db/mysql"

	"github.com/kowiste/boilerplate/pkg/common/controller"
	telemetry "github.com/kowiste/boilerplate/pkg/opentelemetry"
	"github.com/kowiste/boilerplate/pkg/validator"
	conf "github.com/kowiste/boilerplate/src/config"
	ownControl "github.com/kowiste/boilerplate/src/controller"
	"github.com/kowiste/boilerplate/src/controller/grpc"
	"github.com/kowiste/boilerplate/src/controller/rest"
	"github.com/kowiste/boilerplate/src/db"
	assetserv "github.com/kowiste/boilerplate/src/service/asset"
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
	db.New(mysql.New())
	database, err := db.Get()
	if err != nil {
		return
	}
	err = database.Init()
	if err != nil {
		return
	}

	//Init logging
	tp, err := telemetry.New()
	if err != nil {
		return
	}
	//Init services
	assetserv.New(assetserv.WithDatabase(database))
	userservice.New(database)
	tracer := tp.Tracer(cnf.ServiceName)
	//Init controllers
	ctr := make([]ownControl.IController, 0)
	for i := range cnf.Controllers {
		switch cnf.Controllers[i] {
		case controller.Rest:
			ctr = append(ctr, rest.New(rest.WithTracer(&tracer)))
		case controller.GRPC:
			ctr = append(ctr, grpc.New(grpc.WithTracer(&tracer)))
		}
		err = ctr[i].Init()
		if err != nil {
			return
		}
	}

	return
}
