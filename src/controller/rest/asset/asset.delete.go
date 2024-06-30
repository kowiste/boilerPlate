package assetapi

import (
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) deleteAsset(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	err := a.service.Delete(c.Request.Context(), asset.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, asset)
}
