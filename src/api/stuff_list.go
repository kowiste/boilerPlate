package controller

import (
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

// Test App Create Stuff
// @Summary Back Office User
// @Description Create a stuff for the test app
// @Tags Test app stuff
// @Accept json
// @Produce json
// @Param user body model.Stuff true "Stuff data"
// @Success 200 {object} string
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/create [POST]
// @Security Bearer
func (c controller) List(ctx *gin.Context) {
	list := []model.Stuff{}
	c.findAllCore(ctx, &model.Stuff{}, &list)
}
