package grpc

import (
	"context"

	pbAsset "github.com/kowiste/boilerplate/pkg/proto/asset"
)

func (a *GRPC) GetAllAssets(ctx context.Context, req *pbAsset.GetAllAssetsRequest) (*pbAsset.GetAllAssetsResponse, error) {
	assets, err := a.serviceAsset.Assets(ctx)
	if err != nil {
		return nil, err
	}
	return &pbAsset.GetAllAssetsResponse{Assets: assets.ToGRPC()}, nil
}

func (a *GRPC) GetAssetById(ctx context.Context, req *pbAsset.GetByIdRequest) (*pbAsset.GetAssetByIdResponse, error) {
	asset, err := a.serviceAsset.AssetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pbAsset.GetAssetByIdResponse{Asset: asset.ToGRPC()}, nil
}
