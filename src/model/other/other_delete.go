package other

import (
	"github.com/gin-gonic/gin"
)

// Test App Delete Other
// @Summary Delete Other
// @Description Delete other for the test app
// @Tags Other
// @Accept json
// @Produce json
// @Param id path string true "Other ID"
// @Success 200 
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /other/delete [DELETE]
// @Security Bearer
func (s Other) Delete(ctx *gin.Context) {
	s.controller.DeleteCore(ctx, &s)
}
