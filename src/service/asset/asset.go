package assetservice

import (
	"boiler/pkg/errors"
	"boiler/src/model/asset"
	"boiler/src/repository"
	"sync"
)

type AssetService struct {
	asset *asset.Asset
	db    repository.IRepository
}

var (
	instance *AssetService
	once     sync.Once
)

func New(db repository.IRepository) (serv *AssetService, err error) {

	once.Do(func() {

		instance = &AssetService{
			asset: new(asset.Asset),
			db:    db,
		}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func Get() (*AssetService, error) {
	if instance == nil {
		return nil, errors.New("AssetService not set", errors.EErrorServerInternal)
	}
	return instance, nil
}

func (serv *AssetService) GetAsset() *asset.Asset {
	return serv.asset
}
