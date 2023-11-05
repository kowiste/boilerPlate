package controller

import (
	"net/http"
	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

// deleteCore
func (c Controller) DeleteCore(ctx *gin.Context, modelType model.ModelI) {
	err := modelType.SetID(ctx.Param("id"))
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}
	status, err := modelType.OnDelete()
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
		return
	}
	status, err = c.db.Delete(modelType)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
		return
	}
	ctx.Status(status)
}

// AsyncDeleteCore
func (c Controller) AsyncDeleteCore(data model.ModelI) (status int, err error) {
	status, err = data.OnDelete()
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		return
	}
	status, err = c.db.Delete(data)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		return
	}
	return
}
