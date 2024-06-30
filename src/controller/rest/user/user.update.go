package userapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) updateUser(c *gin.Context) {
	user := a.service.GetUser()
	err := c.ShouldBind(&user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	user.ID = c.Param("id")
	err = a.service.Update(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
