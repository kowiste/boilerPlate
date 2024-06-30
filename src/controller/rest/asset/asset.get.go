package assetapi

import (
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) getAssets(c *gin.Context) {
	assets, err := a.service.Get(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, assets)
}

func (a AssetAPI) getAssetByID(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	asset, err := a.service.GetByID(c.Request.Context(), asset.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, asset)
}
