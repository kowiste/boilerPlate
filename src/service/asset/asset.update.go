package assetservice

import (
	"context"
)

// Update modifies an existing asset in the database.
// Parameters:
// - c: The context for the operation.
// Returns:
// - err: An error if the update fails.
func (serv *AssetService) Update(c context.Context) (err error) {
	return serv.db.UpdateAsset(c, serv.asset)
}
