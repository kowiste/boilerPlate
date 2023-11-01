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
	}
	status, err := c.db.Delete(modelType)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		ctx.Status(status)
	}
	ctx.Status(status)
}
