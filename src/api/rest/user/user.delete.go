package userapi

import (
	"boiler/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) deleteUser(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	err := user.Create(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, user)
}
