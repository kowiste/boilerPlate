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
	err := c.db.Create(data)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusCreated, data)
}
