package assetservice

import (
	"context"
)

func (serv *AssetService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteAsset(c, id)
}
