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
	Create(*gin.Context, model.ModelI)
	FindOne(*gin.Context, model.ModelI)
	FindAll(*gin.Context, model.FindAllRequest, model.ModelI, any)
	Update(*gin.Context, model.ModelI, map[string]any)
	Delete(*gin.Context, model.ModelI)
}

func New(db DatabaseI, models ...model.ModelI) {
	c := &Controller{
		api: gin.New(),
		db:  db,
	}
	for _, model := range models {
		model.SetController(c)
		model.InjectAPI()
	}
	c.run()
}

func (c *Controller) run() {
	//gin.SetMode(gin.ReleaseMode)
	c.api.Use(c.timeOut)
	docs.SwaggerInfo.BasePath = "/api"
	c.api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	c.api.Run("0.0.0.0:" + config.Get().Port)
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
