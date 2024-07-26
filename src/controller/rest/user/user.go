package userapi

import (
	userservice "github.com/kowiste/boilerplatesrc/service/user"

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
		userIDGroup := userGroup.Group(":id")
		{
			userIDGroup.GET("", a.getUserByID)
			userIDGroup.PUT("", a.updateUser)
			userIDGroup.DELETE("", a.deleteUser)
		}

	}
}
