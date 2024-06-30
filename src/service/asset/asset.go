package assetservice

import (
	"boiler/src/model/asset"
	"boiler/src/repository"
	"context"
)

type AssetService struct {
	asset *asset.Asset
	db    repository.IRepository
}

func New() (serv *AssetService, err error) {
	database, err := repository.Get()
	if err != nil {
		return
	}
	return &AssetService{
		asset: new(asset.Asset),
		db:    database,
	}, nil
}

func (serv AssetService) Create(c context.Context, asset *asset.Asset) (id string, err error) {
	serv.asset.Validate(c)
	return serv.db.CreateAsset(c, asset)
}

func (serv AssetService) Get(c context.Context) (users []asset.Asset, err error) {
	return serv.db.GetAssets(c)
}

func (serv AssetService) GetByID(c context.Context, id string) (users *asset.Asset, err error) {
	return serv.db.GetAssetByID(c, id)
}

func (serv AssetService) Update(c context.Context, asset *asset.Asset) (err error) {
	return serv.db.UpdateAsset(c, asset)
}
func (serv AssetService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteAsset(c, id)
}
