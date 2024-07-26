package assetservice

import (
	"context"

	"github.com/kowiste/boilerplatesrc/model/asset"
)

// Assets retrieves a list of assets from the database.
// Parameters:
// - c: The context for the operation.
// Returns:
// - users: A list of assets.
// - err: An error if the retrieval fails.
func (serv *AssetService) Assets(c context.Context) (users asset.Assets, err error) {
	return serv.db.Assets(c)
}

// AssetByID retrieves an asset from the database by its ID.
// Parameters:
// - c: The context for the operation.
// - id: The asset ID to find.
// Returns:
// - users: The asset with the specified ID.
// - err: An error if the retrieval fails.
func (serv *AssetService) AssetByID(c context.Context, id string) (users *asset.Asset, err error) {
	return serv.db.AssetByID(c, id)
}
