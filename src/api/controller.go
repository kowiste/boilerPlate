package controller

import (
	"context"
	"errors"
	"net/http"
	"serviceX/docs"
	"serviceX/src/config"
	"serviceX/src/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type controller struct {
	service Service
	db      Database
	engine  *gin.Engine
}

type Service interface {
	Authorization(c *gin.Context)
	Create(c *gin.Context)
	List(c *gin.Context)
	Find(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
type Database interface {
	Create(*gin.Context, model.ModelInterface)
	FindOne(*gin.Context, model.ModelInterface)
	FindAll(*gin.Context, model.FindAllRequest, model.ModelInterface, any)
	Update(*gin.Context, model.ModelInterface, map[string]any)
	Delete(*gin.Context, model.ModelInterface)
}

func New(db Database) *controller {
	c := &controller{
		//service: service,
		engine: gin.New(),
		db:     db,
	}

	api := c.engine.Group("api")
	{
		api.Use(c.timeOut)
		stuff := api.Group("stuff")
		{
			stuff.POST("create", c.Create)
			stuff.GET("list", c.List)
			stuff.GET("find/:id", c.Find)
			stuff.PUT("update/:id", c.Update)
			stuff.DELETE("delete/:id", c.Delete)
		}
	}
	docs.SwaggerInfo.BasePath = "/api"
	c.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return c
}

func (c *controller) Run() {
	c.engine.Run("0.0.0.0:" + config.Get().Port)
}

func (c controller) timeOut(ctx *gin.Context) {
	contx, cancel := context.WithTimeout(ctx.Request.Context(), time.Duration(config.Get().APITimeOut))
	defer cancel()
	// Set the context on the request so that it can be
	// cancelled by the middleware if the timeout is reached.
	ctx.Request = ctx.Request.WithContext(contx)
	ctx.Next()
}

func (c controller) validate(err error) (bool, int, map[string]string) {
	errorMessages := make(map[string]string)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, errorMessage := range validationErrors {
				errorMessages[errorMessage.Namespace()] = errorMessage.Error()
			}
			return false, http.StatusUnprocessableEntity, errorMessages
		} else {
			println(err.Error())
			return false, http.StatusBadRequest, errorMessages
		}
	}
	return true, http.StatusOK, errorMessages
}
