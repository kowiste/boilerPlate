package controller

import (
	"net/http"

	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

// createCore
func (c Controller) CreateCore(ctx *gin.Context, data model.ModelI) {
	valid, code, errorMessages := c.validate(ctx.ShouldBindJSON(data))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}

	if ok, messages := data.CreateValidation(); !ok {
		ctx.JSON(http.StatusUnprocessableEntity, messages)
		return
	}
	status, err := data.OnCreate()
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
		return
	}
	err = c.db.Create(data)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

// AsyncCreateCore
func (c Controller) AsyncCreateCore(data model.ModelI) (status int, err error) {
	status, err = data.OnCreate()
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		return status, err
	}
	err = c.db.Create(data)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
