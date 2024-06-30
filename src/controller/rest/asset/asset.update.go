package assetapi

import (
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) updateAsset(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	err := a.service.Update(c, asset)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, asset)
}