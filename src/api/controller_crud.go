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
func (c controller) createCore(ctx *gin.Context, data model.ModelInterface) {
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
func (c controller) findAllCore(ctx *gin.Context, modelType model.ModelInterface, data any) {
	//Get limit and offset of the request
	request := model.FindAllRequest{
		Limit: 10,
	}
	valid, code, errorMessages := c.validate(ctx.ShouldBindWith(&request, binding.Query))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}
	c.db.FindAll(ctx, request, modelType, data)
}

func (c controller) updateCore(ctx *gin.Context, modelType model.ModelInterface) {
	// Schema validation
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
	}
	ctx.Request.Body = io.NopCloser(bytes.NewReader(data))

	updateData := make(map[string]interface{})
	err = json.Unmarshal(data, &updateData)
	if err != nil {
	}
	// Validation if data is proper
	modelType.BeforeValidation()
	valid, code, errorMessages := c.validate(ctx.ShouldBindWith(&modelType, binding.JSON))
	if !valid {
		ctx.JSON(code, errorMessages)
		return
	}

	if ok, messages := modelType.UpdateValidation(); !ok {
		ctx.JSON(422, messages)
		return
	}
	modelType.AfterValidation()
	
	c.db.Update(ctx, modelType, updateData)
}

func (c controller) deleteCore(ctx *gin.Context, modelType model.ModelInterface) {

}
