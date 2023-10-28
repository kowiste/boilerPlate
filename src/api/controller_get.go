package controller

import (
	"net/http"
	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// FindOne
func (c Controller) FindOne(ctx *gin.Context, modelType model.ModelI) {
	err := modelType.SetID(ctx.Param("id"))
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(http.StatusBadRequest)
	}
	status, err := c.db.FindOne(modelType)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
		return
	}
	ctx.JSON(status, modelType)
}

// findAllCore
func (c Controller) FindAllCore(ctx *gin.Context, modelType model.ModelI, data any) {
	request := model.FindAllRequest{ //Get limit and offset of the request
		Limit: 10,
	}
	valid, code, errorMessages := c.validate(ctx.ShouldBindWith(&request, binding.Query))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}
	filters := ctx.QueryMap("filter")

	status, count, err := c.db.FindAll(filters, request, modelType, data)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
		return
	}
	ctx.JSON(status, model.FindAllResponse{
		Count: count,
		Data:  data,
	})
}
