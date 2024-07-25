package assetservice

import (
	"context"
)

func (serv *AssetService) Update(c context.Context) (err error) {
	return serv.db.UpdateAsset(c, serv.asset)
}
