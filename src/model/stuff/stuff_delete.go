package stuff

import (
	"github.com/gin-gonic/gin"
)

// Test App Delete Stuff
// @Summary Delete Stuff
// @Description Delete Stuff for the test app
// @Tags Stuff
// @Accept json
// @Produce json
// @Param id path string true "Stuff ID"
// @Success 200 
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/delete [DELETE]
// @Security Bearer
func (s Stuff) Delete(ctx *gin.Context) {
	s.controller.DeleteCore(ctx, &s)
}
