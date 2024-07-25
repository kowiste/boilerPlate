package userapi

import (
	userservice "boiler/src/service/user"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	service *userservice.UserService
}

func New() (api *UserAPI, err error) {
	s, err := userservice.Get()
	api = &UserAPI{
		service: s,
	}
	return
}

func (a *UserAPI) Routes(r *gin.Engine) {
	r.GET("/users", a.getUsers)
	userGroup := r.Group("user")
	{
		userGroup.POST("", a.createUser)
		userGroup.GET(":id", a.getUserByID)
		userGroup.PUT(":id", a.updateUser)
		userGroup.DELETE(":id", a.deleteUser)
	}
}
