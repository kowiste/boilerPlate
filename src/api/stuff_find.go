package controller

import (
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

// Test App Find Stuff
// @Summary Test App Find Stuff
// @Description Find one stuff for the test app
// @Tags Test app stuff
// @Accept json
// @Produce json
// @Param user body model.Stuff true "Stuff data"
// @Success 200 {object} string
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/{id} [GET]
// @Security Bearer
func (c controller) Find(ctx *gin.Context) {
	var data model.Stuff
	c.db.FindOne(ctx, &data)
}
