package assetapi

import (
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) getAssets(c *gin.Context) {
	assets := new(asset.Assets)
	err := assets.Get(c.Request.Context(), nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, assets)
}

func (a AssetAPI) getAssetByID(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	err := asset.Get(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, asset)
}
