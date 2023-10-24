package model

import "github.com/gin-gonic/gin"

type ControllerI interface {
	GetAPI() *gin.Engine
	CreateCore(ctx *gin.Context, data ModelI)
	FindOne(ctx *gin.Context, modelType ModelI)
	FindAllCore(ctx *gin.Context, modelType ModelI, data any)
	UpdateCore(ctx *gin.Context, modelType ModelI)
	DeleteCore(ctx *gin.Context, modelType ModelI)
}
