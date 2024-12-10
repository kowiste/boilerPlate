package assethandler

import (
	"ddd/internal/features/asset/domain"
)

type CreateAssetRequest struct {
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description"`
	Type        string                 `json:"type" binding:"required"`
	Properties  map[string]interface{} `json:"properties"`
}

type UpdateAssetRequest struct {
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

type AssetResponse struct {
	ID          string `json:"id"`
	OrgID       string `json:"orgId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func ToAssetResponse(a *domain.Asset) AssetResponse {
	return AssetResponse{
		ID:          a.ID(),
		OrgID:       a.OrgID(),
		Name:        a.Name(),
		Description: a.Description(),
		UpdatedAt:   a.UpdatedAt().Unix(),
	}
}

func ToAssetResponses(assets []*domain.Asset) []AssetResponse {
	responses := make([]AssetResponse, len(assets))
	for i, a := range assets {
		responses[i] = ToAssetResponse(a)
	}
	return responses
}
