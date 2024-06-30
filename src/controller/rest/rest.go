package rest

import (
	conf "boiler/src/config"
	assetapi "boiler/src/controller/rest/asset"
	userapi "boiler/src/controller/rest/user"

	"github.com/gin-gonic/gin"
	"github.com/kowiste/config"
)

type API struct {
	router *gin.Engine
}

func New() (api *API) {
	return &API{}
}
func (a *API) Init() (err error) {
	c, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		return
	}
	a.router = gin.Default()
	user, err := userapi.New()
	if err != nil {
		return
	}
	user.Routes(a.router)
	asset, err := assetapi.New()
	if err != nil {
		return
	}
	asset.Routes(a.router)

	go func() {
		err = a.router.Run(":" + c.ServicePort) // listen and serve on port 8080
		if err != nil {
			panic(err)
		}
	}()
	return
}
