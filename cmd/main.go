package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"test.com/config"
	"test.com/model"
	"test.com/service"
	"test.com/handler/validator"
)

type Service interface {
	Authorization(c *gin.Context)
	Create(c *gin.Context)
	List(c *gin.Context)
	Find(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// @title           Test app API
// @version         1.0
// @description     API of test app.
// @BasePath  /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	var core model.Service
	err := config.CreateInstance()
	if err != nil {
		panic(err)
	}
	validator.New()
	err = validator.GetInstance().Validate(config.Get())
	if err != nil {
		panic(err)
	}
	core = service.Init()

	r := gin.Default()
	api := r.Group("api")
	{
		api.Use(core.Authorization)
		stuff := api.Group("stuff")
		{
			stuff.POST("create", core.Create)
			stuff.GET("list", core.List)
			stuff.GET("find/:id", core.Find)
			stuff.PATCH("update/:id", core.Update)
			stuff.DELETE("delete/:id", core.Delete)
		}
	}
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + config.Get().Port)
}
