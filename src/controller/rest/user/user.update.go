package userapi

import (
	"net/http"

	"github.com/kowiste/boilerplate/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) updateUser(c *gin.Context) {
	user := a.service.GetUser()
	err := c.ShouldBind(&user)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	user.ID = c.Param("id")
	err = a.service.Update(c)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.Status(http.StatusOK)
}
