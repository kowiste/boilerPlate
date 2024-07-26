package assetapi

import (
	"net/http"

	"github.com/kowiste/boilerplatepkg/errors"
	"github.com/kowiste/boilerplatesrc/model/asset"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) deleteAsset(c *gin.Context) {
	asset := new(asset.Asset)
	asset.ID = c.Param("id")
	err := a.service.Delete(c.Request.Context(), asset.ID)
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, asset)
}
