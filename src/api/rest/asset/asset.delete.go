package assetapi

import (
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) deleteAsset(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	err := asset.Create(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, asset)
}
