package stuff

import (
	"github.com/gin-gonic/gin"
)

// Test App Find All Stuff
// @Summary Find All Stuff
// @Description Return array of Stuff for the test app
// @Tags Stuff
// @Produce json
// @Param filter[name] query string false "Filter by name"
// @Success 200 {object} model.FindAllResponse{data=[]stuff.Stuff{id=string}}
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/list [GET]
// @Security Bearer
func (s Stuff) List(ctx *gin.Context) {
	list := []Stuff{}
	s.controller.FindAllCore(ctx, &s, &list)
}
