package userapi

import (
	"boiler/pkg/errors"
	"boiler/src/model/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a UserAPI) userQuery(c *gin.Context) (input *user.FindUsersInput, err error) {
	input = new(user.FindUsersInput)
	input.Text = c.Query("text")
	ageString := c.Query("age")

	age, err := strconv.Atoi(ageString)
	if err != nil {
		return nil, errors.New("", errors.EErrorUnhandled)
	}
	input.Age = age
	return

}
func (a UserAPI) getUsers(c *gin.Context) {
	input, err := a.userQuery(c)
	if err != nil {
		errors.RestError(c.Writer, err)
	}
	users, err := a.service.Users(c, input)
	if err != nil {
		errors.RestError(c.Writer, err)
	}
	c.JSON(http.StatusOK, users)
}

func (a UserAPI) getUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := a.service.UserByID(c, id)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
