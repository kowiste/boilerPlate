package userapi

import (
	"boiler/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) getUsers(c *gin.Context) {
	users, err := a.service.Users(c)
	if err != nil {
		errors.RestError(c.Writer, err)
	}
	c.JSON(http.StatusOK, users)
}

func (a UserAPI) getUserByID(c *gin.Context) {
	user := a.service.GetUser()
	user.ID = c.Param("id")
	user, err := a.service.UserByID(c)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
