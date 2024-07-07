package assetapi

import (
	"boiler/pkg/errors"
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) getAssets(c *gin.Context) {
	assets, err := a.service.Assets(c.Request.Context())
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, assets)
}

func (a AssetAPI) getAssetByID(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	asset, err := a.service.AssetByID(c.Request.Context(), asset.ID)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, asset)
}
