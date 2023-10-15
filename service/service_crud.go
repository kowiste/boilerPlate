package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"test.com/model"
)

// createCore
func (s service) createCore(c *gin.Context, data model.ModelInterface) {
	valid, code, errorMessages := s.validate(c.ShouldBindJSON(data))
	if !valid {
		c.JSON(code, errorMessages)
		return
	}

	if ok, messages := data.CreateValidation(); !ok {
		c.JSON(http.StatusUnprocessableEntity, messages)
		return
	}
	s.db.Create(c, data)
}

// findAllCore
func (s service) findAllCore(c *gin.Context, modelType model.ModelInterface, data any) {
	//Get limit and offset of the request
	request := model.FindAllRequest{
		Limit: 10,
	}
	valid, code, errorMessages := s.validate(c.ShouldBindWith(&request, binding.Query))
	if !valid {
		c.JSON(code, errorMessages)
		return
	}
	s.db.FindAll(c, request, modelType, data)
}

func (s service) updateCore(c *gin.Context, modelType model.ModelInterface) {
	// Schema validation
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(data))

	updateData := make(map[string]interface{})
	err = json.Unmarshal(data, &updateData)
	if err != nil {
	}
	// Validation if data is proper
	modelType.BeforeValidation()
	valid, code, errorMessages := s.validate(c.ShouldBindWith(&modelType, binding.JSON))
	if !valid {
		c.JSON(code, errorMessages)
		return
	}

	if ok, messages := modelType.UpdateValidation(); !ok {
		c.JSON(422, messages)
		return
	}
	modelType.AfterValidation()
	s.db.Update(c, modelType, updateData)
}

func (s service) deleteCore(c *gin.Context, modelType model.ModelInterface) {

}
