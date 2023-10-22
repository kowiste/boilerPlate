package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"serviceX/src/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// createCore
func (c Controller) CreateCore(ctx *gin.Context, data model.ModelInterface) {
	valid, code, errorMessages := c.validate(ctx.ShouldBindJSON(data))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}

	if ok, messages := data.CreateValidation(); !ok {
		ctx.JSON(http.StatusUnprocessableEntity, messages)
		return
	}
	c.db.Create(ctx, data)
}

// findAllCore
func (c Controller) FindOne(ctx *gin.Context, modelType model.ModelInterface) {
	c.db.FindOne(ctx, modelType)
}

// findAllCore
func (c Controller) FindAllCore(ctx *gin.Context, modelType model.ModelInterface, data any) {
	request := model.FindAllRequest{ //Get limit and offset of the request
		Limit: 10,
	}
	valid, code, errorMessages := c.validate(ctx.ShouldBindWith(&request, binding.Query))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}
	c.db.FindAll(ctx, request, modelType, data)
}

// updateCore
func (c Controller) UpdateCore(ctx *gin.Context, modelType model.ModelInterface) {
	// Schema validation
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Request.Body = io.NopCloser(bytes.NewReader(data))

	updateData := make(map[string]interface{})
	err = json.Unmarshal(data, &updateData)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	modelType.BeforeValidation() //run logic before validate
	valid, code, errorMessages := c.validate(ctx.ShouldBindWith(&modelType, binding.JSON))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}

	if ok, messages := modelType.UpdateValidation(); !ok {
		ctx.JSON(http.StatusUnprocessableEntity, messages)
		return
	}
	modelType.AfterValidation()

	c.db.Update(ctx, modelType, updateData)
}

// deleteCore
func (c Controller) DeleteCore(ctx *gin.Context, modelType model.ModelInterface) {
	c.db.Delete(ctx, modelType)
}
