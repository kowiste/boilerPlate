package stuff

import (
	"github.com/gin-gonic/gin"
)

// Test App Update Stuff
// @Summary Update Stuff
// @Description Update Stuff for the test app
// @Tags Stuff
// @Accept json
// @Produce json
// @Param id path string true "Stuff ID"
// @Param stuff body stuff.Stuff true "Stuff data"
// @Success 200 {object} stuff.Stuff{id=string}
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/update/{id} [PATCH]
// @Security Bearer
func (s Stuff) Update(ctx *gin.Context) {
	s.controller.UpdateCore(ctx, &s)
}
