package assetapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) updateAsset(c *gin.Context) {
	asset := a.service.GetAsset()
	asset.ID = c.Param("id")
	err := a.service.Update(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, asset)
}
