package userapi

import (
	"boiler/src/model/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) createUser(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	c.ShouldBind(&user)

	id, err := a.service.Create(c, user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, id)
}
