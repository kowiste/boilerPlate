package stuff

import (
	"github.com/gin-gonic/gin"
)

// Test App Find Stuff
// @Summary Test App Find One Stuff
// @Description Find one Stuff for the test app
// @Tags Stuff
// @Produce json
// @Param id path string true "Stuff ID"
// @Success 200 {object} stuff.Stuff{id=string}
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/find{id} [GET]
// @Security Bearer
func (s Stuff) Find(ctx *gin.Context) {
	s.controller.FindOne(ctx, &s)
}
