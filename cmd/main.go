package main

import (
	controller "serviceX/src/api"
	"serviceX/src/config"
	"serviceX/src/handler/broker/nats"
	"serviceX/src/handler/database/nosql"
	"serviceX/src/handler/database/sql"
	"serviceX/src/handler/log"
	"serviceX/src/handler/validator"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

// @title           Test app API
// @version         1.0
// @description     API of test app.
// @BasePath  /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	err := config.CreateInstance()
	if err != nil {
		panic(err)
	}
	//Initializate log
	log.CreateInstance(log.ErrorLevel)

	//Initializate validation
	validator.New()
	err = validator.Get().Validate(config.Get())
	if err != nil {
		panic(err)
	}

	//Initializate broker
	err = nats.CreateInstance(config.Get().Name)
	if err != nil {
		panic(err)
	}
	nats.Get().SetMessageEvent(func(msg []byte) error {
		return nil
	})
	log.Get().SetLocal(true)
	log.Get().SetChannels(nats.Get().GetChannel())

	//Config database SQL
	db := sql.CreatePostgres(&model.Stuff{})
	defer func() {
		db.Close()
	}()

	//SQL controler
	controllerSQL := controller.New(db)
	gin.SetMode(gin.ReleaseMode)
	go controllerSQL.Run()

	dbMongo := nosql.CreateMongo(config.Get().Name)
	defer func() {
		dbMongo.Close()
	}()

	//SQL controler
	controllerMongo := controller.New(dbMongo)
	gin.SetMode(gin.ReleaseMode)
	controllerMongo.Run()
}
