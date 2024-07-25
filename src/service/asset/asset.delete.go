package assetservice

import (
	"context"
)

// Delete removes an asset from the database by ID.
// Parameters:
// - c: The context for the operation.
// - id: The asset ID to delete.
// Returns:
// - err: An error if the deletion fails.
func (serv *AssetService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteAsset(c, id)
}
