package assetapi

import (
	"net/http"

	"github.com/kowiste/boilerplatepkg/errors"

	"github.com/gin-gonic/gin"
)

func (a AssetAPI) createAsset(c *gin.Context) {
	asset := a.service.GetAsset()
	asset.ID = c.Param("id")
	id, err := a.service.Create(c.Request.Context())
	if err != nil {
		errors.RestError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, id)
}
