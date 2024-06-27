package userapi

import "github.com/gin-gonic/gin"

type UserAPI struct{}

func New() *UserAPI {
	return &UserAPI{}
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
