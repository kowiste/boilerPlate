package main

import (
	controller "serviceX/src/api"
	"serviceX/src/config"
	"serviceX/src/handler/broker/nats"
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

	//Initializate broker
	err = nats.CreateInstance(config.Get().Name)
	if err != nil {
		panic(err)
	}
	nats.Get().SetMessageEvent(func(msg []byte) error {
		return nil
	})

	log.Get().SetChannels(nats.Get().GetChannel())

	//Config database
	db := sql.CreatePostgres(&model.Stuff{})
	//db := nosql.CreateMongo("service1")
	defer func() {
		db.Close()
	}()
	validator.New()
	err = validator.Get().Validate(config.Get())
	if err != nil {
		panic(err)
	}
	log.Get().Print(log.InfoLevel, "hi")

	controller := controller.New(db)
	gin.SetMode(gin.ReleaseMode)
	controller.Run()
}
