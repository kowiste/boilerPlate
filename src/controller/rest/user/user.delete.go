package userapi

import (
	"boiler/pkg/errors"
	"boiler/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) deleteUser(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	err := a.service.Delete(c, user.ID)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.Status(http.StatusOK)
}
