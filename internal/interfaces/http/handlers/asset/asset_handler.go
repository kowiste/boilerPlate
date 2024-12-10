package assethandler

import (
	"net/http"

	"ddd/internal/features/asset/app"
	"ddd/internal/features/asset/domain"
	"ddd/shared/httputil"
	"ddd/shared/logger"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	logger       logger.Logger
	assetService app.AssetService
}

type Dependencies struct {
	Logger       logger.Logger
	AssetService app.AssetService
}

func New(deps Dependencies) *AssetHandler {
	return &AssetHandler{
		logger:       deps.Logger,
		assetService: deps.AssetService,
	}
}

// @Summary Create a new asset
// @Description Create a new asset for the organization
// @Tags assets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param orgid path string true "Organization ID" example:"org123"
// @Param asset body CreateAssetRequest true "Asset creation request"
// @Success 201 {object} domain.Asset
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/{orgid}/assets [post]
func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var req CreateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orgID := httputil.GetOrgID(c)
	if orgID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization ID is required"})
		return
	}

	cmd := app.CreateAssetCommand{
		OrgID:       orgID,
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		Properties:  req.Properties,
	}

	result, err := h.assetService.CreateAsset(c.Request.Context(), cmd)
	if err != nil {
		h.logger.Error(c.Request.Context(), err, "Failed to create asset", map[string]interface{}{
			"error": err.Error(),
			"orgID": orgID,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset"})
		return
	}

	c.JSON(http.StatusCreated, ToAssetResponse(result))
}

// @Summary Get an asset by ID
// @Description Get an asset by its ID for the organization
// @Tags assets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param orgid path string true "Organization ID" example:"org123"
// @Param id path string true "Asset ID" example:"550e8400-e29b-41d4-a716-446655440000"
// @Success 200 {object} domain.Asset
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/{orgid}/assets/{id} [get]
func (h *AssetHandler) GetAsset(c *gin.Context) {
	assetID := c.Param("id")
	orgID := httputil.GetOrgID(c)

	if orgID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization ID is required"})
		return
	}

	result, err := h.assetService.GetAsset(c.Request.Context(), orgID, assetID)
	if err != nil {
		if err == domain.ErrAssetNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
			return
		}
		h.logger.Error(c.Request.Context(), err, "Failed to get asset", map[string]interface{}{
			"error":   err.Error(),
			"orgID":   orgID,
			"assetID": assetID,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get asset"})
		return
	}

	c.JSON(http.StatusOK, ToAssetResponse(result))
}

// @Summary List all assets
// @Description List all assets for the organization
// @Tags assets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param orgid path string true "Organization ID" example:"org123"
// @Success 200 {object} struct{assets []domain.Asset} "Array of assets wrapped in assets field"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/{orgid}/assets [get]
func (h *AssetHandler) ListAssets(c *gin.Context) {
	orgID := httputil.GetOrgID(c)
	if orgID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization ID is required"})
		return
	}

	results, err := h.assetService.ListAssets(c.Request.Context(), orgID)
	if err != nil {
		// h.logger.Error(c.Request.Context(), "Failed to list assets", map[string]interface{}{
		// 	"error": err.Error(),
		// 	"orgID": orgID,
		// })
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list assets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"assets": ToAssetResponses(results)})
}

// @Summary Update an asset
// @Description Update an existing asset by ID
// @Tags assets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param orgid path string true "Organization ID" example:"org123"
// @Param id path string true "Asset ID" example:"550e8400-e29b-41d4-a716-446655440000"
// @Param asset body UpdateAssetRequest true "Asset update request"
// @Success 200 {object} domain.Asset
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/{orgid}/assets/{id} [put]

func (h *AssetHandler) UpdateAsset(c *gin.Context) {
	assetID := c.Param("id")
	orgID := c.GetString("orgID")

	var req UpdateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := app.UpdateAssetCommand{
		ID:          assetID,
		OrgID:       orgID,
		Name:        req.Name,
		Description: req.Description,
		Properties:  req.Properties,
	}

	result, err := h.assetService.UpdateAsset(c.Request.Context(), cmd)
	if err != nil {
		if err == domain.ErrAssetNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
			return
		}
		// h.logger.Error(c.Request.Context(), "Failed to update asset", map[string]interface{}{
		// 	"error":   err.Error(),
		// 	"orgID":   orgID,
		// 	"assetID": assetID,
		// })
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	c.JSON(http.StatusOK, ToAssetResponse(result))
}

// @Summary Delete an asset
// @Description Delete an asset by ID
// @Tags assets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param orgid path string true "Organization ID" example:"org123"
// @Param id path string true "Asset ID" example:"550e8400-e29b-41d4-a716-446655440000"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/{orgid}/assets/{id} [delete]

func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	assetID := c.Param("id")
	orgID := c.GetString("orgID")

	if orgID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization ID is required"})
		return
	}

	err := h.assetService.DeleteAsset(c.Request.Context(), orgID, assetID)
	if err != nil {
		if err == domain.ErrAssetNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
			return
		}
		// h.logger.Error(c.Request.Context(), "Failed to delete asset", map[string]interface{}{
		// 	"error":   err.Error(),
		// 	"orgID":   orgID,
		// 	"assetID": assetID,
		// })
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset"})
		return
	}

	c.Status(http.StatusNoContent)
}
