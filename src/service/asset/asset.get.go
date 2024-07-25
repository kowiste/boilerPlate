package assetservice

import (
	"boiler/src/model/asset"
	"context"
)

func (serv *AssetService) Assets(c context.Context) (users []asset.Asset, err error) {
	return serv.db.Assets(c)
}

func (serv *AssetService) AssetByID(c context.Context, id string) (users *asset.Asset, err error) {
	return serv.db.AssetByID(c, id)
}
