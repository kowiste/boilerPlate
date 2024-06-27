package rest

import (
	assetapi "boiler/src/api/rest/asset"
	userapi "boiler/src/api/rest/user"
	conf "boiler/src/config"
	"fmt"

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
		fmt.Println("Error getting config:", err)
		return
	}
	a.router = gin.Default()
	user := userapi.New()
	user.Routes(a.router)
	asset := assetapi.New()
	asset.Routes(a.router)

	go func() {
		err = a.router.Run(":" + c.ServicePort) // listen and serve on port 8080
		if err != nil {
			panic(err)
		}
	}()
	return
}
