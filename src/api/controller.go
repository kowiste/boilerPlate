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

type Controller struct {
	db     Database
	engine *gin.Engine
}
type Database interface {
	Create(*gin.Context, model.ModelInterface)
	FindOne(*gin.Context, model.ModelInterface)
	FindAll(*gin.Context, model.FindAllRequest, model.ModelInterface, any)
	Update(*gin.Context, model.ModelInterface, map[string]any)
	Delete(*gin.Context, model.ModelInterface)
}

func New(db Database) *Controller {
	c := &Controller{
		engine: gin.New(),
		db:     db,
	}

	c.engine.Use(c.timeOut)
	docs.SwaggerInfo.BasePath = "/api"
	c.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return c
}

func (c *Controller) Run() {
	c.engine.Run("0.0.0.0:" + config.Get().Port)
}

func (c *Controller) GetEngine() *gin.Engine {
	return c.engine
}

func (c Controller) timeOut(ctx *gin.Context) {
	contx, cancel := context.WithTimeout(ctx.Request.Context(), time.Duration(config.Get().APITimeOut))
	defer cancel()
	ctx.Request = ctx.Request.WithContext(contx)
	ctx.Next()
}

func (c Controller) validate(err error) (bool, int, map[string]string) {
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
