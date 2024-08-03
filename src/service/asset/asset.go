package assetservice

import (
	"sync"

	"github.com/Kowiste/boilerPlate/src/repository"
	"github.com/kowiste/boilerplate/pkg/errors"
	"github.com/kowiste/boilerplate/src/model/asset"
)

type AssetService struct {
	asset *asset.Asset
	db    repository.IRepository
}

var (
	instance *AssetService
	once     sync.Once
)

func New(db repository.IRepository) (serv *AssetService) {

	once.Do(func() {

		instance = &AssetService{
			asset: new(asset.Asset),
			db:    db,
		}
	})

	return instance
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
