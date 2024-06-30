package userapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) createUser(c *gin.Context) {
	user := a.service.GetUser()
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := a.service.Create(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, id)
}
