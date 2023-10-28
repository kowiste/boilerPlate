package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	ErrValidation string = "validation error"
)

// updateCore
func (c Controller) UpdateCore(ctx *gin.Context, modelType model.ModelI) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	updateData := make(map[string]interface{})
	err = json.Unmarshal(data, &updateData)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}
	modelType.SetID(ctx.Param("id"))

	//Validation zone
	modelType.BeforeValidation() //run logic before validate
	valid, code, errorMessages := c.validate(ctx.ShouldBindWith(&modelType, binding.JSON))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}

	if ok, messages := modelType.UpdateValidation(); !ok {
		log.Get().Print(log.ErrorLevel, ErrValidation)
		ctx.JSON(http.StatusUnprocessableEntity, messages)
		return
	}
	modelType.AfterValidation() //run logic after validate

	//Database
	status, err := c.db.Update(modelType, updateData)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
		return
	}
	ctx.JSON(status, modelType)
}
