package other

import (
	"github.com/gin-gonic/gin"
)

// Test App Find Other
// @Summary Test App Find One Other
// @Description Find one other for the test app
// @Tags Other
// @Produce json
// @Param id path string true "Other ID"
// @Success 200 {object} other.Other{id=string}
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /other/find/{id} [GET]
// @Security Bearer
func (s Other) Find(ctx *gin.Context) {
	s.controller.FindOne(ctx, &s)
}
