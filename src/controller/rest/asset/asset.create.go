package assetapi

import (
	"boiler/src/model/asset"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) createAsset(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	id, err := a.service.Create(c.Request.Context(), asset)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, id)
}