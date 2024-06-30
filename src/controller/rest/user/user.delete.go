package userapi

import (
	"boiler/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) deleteUser(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	err := a.service.Delete(c,user.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
