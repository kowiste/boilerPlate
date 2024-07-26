package userapi

import (
	"net/http"

	"github.com/kowiste/boilerplatepkg/errors"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) deleteUser(c *gin.Context) {
	id := c.Param("id")

	err := a.service.Delete(c, id)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.Status(http.StatusOK)
}
