package other

import (
	"github.com/gin-gonic/gin"
)

// Test App Update other
// @Summary Update Other
// @Description Update other for the test app
// @Tags Other
// @Accept json
// @Produce json
// @Param id path string true "Other ID"
// @Param other body other.Other true "Other data"
// @Success 200 {object} other.Other{id=string}
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /other/update/{id} [PATCH]
// @Security Bearer
func (s Other) Update(ctx *gin.Context) {
	s.controller.UpdateCore(ctx, &s)
}
