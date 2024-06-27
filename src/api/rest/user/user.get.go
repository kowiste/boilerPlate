package userapi

import (
	"boiler/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) getUsers(c *gin.Context) {
	users := new(user.Users)
	err := users.Get(c.Request.Context(), nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, users)
}

func (a UserAPI) getUserByID(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	err := user.Get(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, user)
}
