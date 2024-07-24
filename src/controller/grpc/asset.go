package grpc

import (
	pbAsset "boiler/doc/proto/asset"
	"context"
)

func (a *GRPC) GetAllAssets(ctx context.Context, req *pbAsset.GetAllAssetsRequest) (*pbAsset.GetAllAssetsResponse, error) {
	//assets, err := a.serviceAsset.GetAsset(req.ParentId)
	assets := a.serviceAsset.GetAsset()

	return &pbAsset.GetAllAssetsResponse{Assets: assets}, nil
}

func (a *GRPC) GetAssetById(ctx context.Context, req *pbAsset.GetByIdRequest) (*pbAsset.GetAssetByIdResponse, error) {
	asset, err := a.serviceAsset.AssetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pbAsset.GetAssetByIdResponse{Asset: asset}, nil
}
