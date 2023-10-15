package main

import (
	controller "serviceX/src/api"
	"serviceX/src/config"
	"serviceX/src/handler/database/sql"
	"serviceX/src/handler/validator"

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
	db := sql.CreateInstance()
	defer func() {
		db.Close()
	}()
	validator.New()
	err = validator.Get().Validate(config.Get())
	if err != nil {
		panic(err)
	}
	controller := controller.New()

	gin.SetMode(gin.ReleaseMode)
	controller.Run()
}
