package userapi

import (
	"boiler/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) updateUser(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	err := a.service.Update(c, user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}
