package other

import (
	"github.com/gin-gonic/gin"
)

// Test App Create Other
// @Summary Create other
// @Description Create a other for the test app
// @Tags Other
// @Accept json
// @Produce json
// @Param other body other.Other true "Other data"
// @Success 200 {object} string
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /other/create [POST]
// @Security Bearer
func (s Other) Create(ctx *gin.Context) {
	s.controller.CreateCore(ctx, &s)
}
