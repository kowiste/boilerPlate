package stuff

import (
	"github.com/gin-gonic/gin"
)

// Test App Create Stuff
// @Summary Create Stuff
// @Description Create a stuff for the test app
// @Tags Stuff
// @Accept json
// @Produce json
// @Param stuff body stuff.Stuff true "Stuff data"
// @Success 200 {object} string
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/create [POST]
// @Security Bearer
func (s Stuff) Create(ctx *gin.Context) {
	s.controller.CreateCore(ctx, &s)
}
