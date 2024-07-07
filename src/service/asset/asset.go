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
func (serv *AssetService) GetAsset() *asset.Asset {
	return serv.asset
}
func (serv AssetService) Create(c context.Context) (id string, err error) {
	serv.asset.Validate(c)
	return serv.db.CreateAsset(c, serv.asset)
}

func (serv AssetService) Assets(c context.Context) (users []asset.Asset, err error) {
	return serv.db.Assets(c)
}

func (serv AssetService) AssetByID(c context.Context, id string) (users *asset.Asset, err error) {
	return serv.db.AssetByID(c, id)
}

func (serv AssetService) Update(c context.Context) (err error) {
	return serv.db.UpdateAsset(c, serv.asset)
}
func (serv AssetService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteAsset(c, id)
}
