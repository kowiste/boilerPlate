package assetservice

import (
	"context"
)

func (serv *AssetService) Create(c context.Context) (id string, err error) {
	serv.asset.Validate(c)
	return serv.db.CreateAsset(c, serv.asset)
}
