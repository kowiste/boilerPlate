package asset

import (
	pbAsset "github.com/kowiste/boilerplate/pkg/proto/asset"
)

type Asset struct {
	ID          string `json:"id"`
	ParentID    string `json:"parentID" validate:"uuid"`
	Description string `json:"description"`
}

type Assets []Asset

func (a Asset) TableName() string {
	return "assets"
}
func (a Asset) ToGRPC() *pbAsset.Asset {
	return &pbAsset.Asset{
		Id:          a.ID,
		ParentId:    a.ParentID,
		Description: a.Description,
	}
}

// ToGRPC converts the Assets slice to a Asset protobuf message.
func (a Assets) ToGRPC() []*pbAsset.Asset {
	assets := make([]*pbAsset.Asset, len(a))
	for i, asset := range a {
		assets[i] = asset.ToGRPC()
	}

	return assets

}
