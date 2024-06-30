package userapi

import (
	"boiler/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) getUsers(c *gin.Context) {
	users, err := a.service.Get(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, users)
}

func (a UserAPI) getUserByID(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	user, err := a.service.GetByID(c, user.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, user)
}
