package other

import (
	"github.com/gin-gonic/gin"
)

// Test App Find All Other
// @Summary Find All Other
// @Description Return array of other for the test app
// @Tags Other
// @Produce json
// @Param filter[name] query string false "Filter by name"
// @Success 200 {object} model.FindAllResponse{data=[]other.Other{id=string}}
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /other/list [GET]
// @Security Bearer
func (s Other) List(ctx *gin.Context) {
	list := []Other{}
	s.controller.FindAllCore(ctx, &s, &list)
}
