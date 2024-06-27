package assetapi

import "github.com/gin-gonic/gin"

type AssetAPI struct{}

func New() *AssetAPI {
	return &AssetAPI{}
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
