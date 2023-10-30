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
	db  DatabaseI
	api *gin.Engine
}

type DatabaseI interface {
	Create(data model.ModelI) (err error)
	FindOne(model.ModelI) (status int, err error)
	FindAll(map[string]string, model.FindAllRequest, model.ModelI, any) (status int, count int64, err error)
	Update(model.ModelI, map[string]any) (status int, err error)
	Delete(data model.ModelI) (status int, err error)
}

func New(port string, db DatabaseI, models ...model.ModelI) {
	c := &Controller{
		api: gin.New(),
		db:  db,
	}
	//Loading api
	for _, model := range models {
		model.SetController(c)
		model.InjectAPI()
	}
	c.run(port)
}

func (c *Controller) run(port string) {
	//gin.SetMode(gin.ReleaseMode)
	c.api.Use(c.timeOut)
	docs.SwaggerInfo.BasePath = "/api"
	c.api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	c.api.Run("0.0.0.0:" + port)
}

func (c Controller) GetAPI() *gin.Engine {
	return c.api
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
