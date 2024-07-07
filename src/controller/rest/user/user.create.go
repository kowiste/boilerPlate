package userapi

import (
	"boiler/pkg/errors"
	"boiler/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) createUser(c *gin.Context) {
	user := a.service.GetUser()
	err := c.ShouldBind(&user)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	id, err := a.service.Create(c)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse{
		ID: id,
	})
}
