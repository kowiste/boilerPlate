package assetapi

import (
	assetservice "boiler/src/service/asset"

	"github.com/gin-gonic/gin"
)

type AssetAPI struct {
	service *assetservice.AssetService
}

func New() (api *AssetAPI, err error) {
	s, err := assetservice.Get()
	api = &AssetAPI{
		service: s,
	}
	return
}

func (a *AssetAPI) Routes(r *gin.Engine) {
	r.GET("/assets", a.getAssets)
	assetGroup := r.Group("asset")
	{
		assetGroup.POST("", a.createAsset)
		assetIDGroup := assetGroup.Group(":id")
		{
			assetIDGroup.GET("", a.getAssetByID)
			assetIDGroup.PUT("", a.updateAsset)
			assetIDGroup.DELETE("", a.deleteAsset)
		}

	}
}
