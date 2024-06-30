package assetapi

import (
	assetservice "boiler/src/service/asset"

	"github.com/gin-gonic/gin"
)

type AssetAPI struct {
	service *assetservice.AssetService
}

func New() (api *AssetAPI, err error) {
	s, err := assetservice.New()
	api = &AssetAPI{
		service: s,
	}
	return
}

func (a *AssetAPI) Routes(r *gin.Engine) {
	r.GET("/assets", a.getAssets)
	userGroup := r.Group("asset")
	{
		userGroup.POST("", a.createAsset)
		userGroup.GET(":id", a.getAssetByID)
		userGroup.PUT(":id", a.updateAsset)
		userGroup.DELETE(":id", a.deleteAsset)
	}
}
