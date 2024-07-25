package assetservice

import (
	"context"
)

// Create adds a new asset to the database.
// Parameters:
// - c: The context for the operation.
// Returns:
// - id: The newly created asset's ID.
// - err: An error if the creation fails.
func (serv *AssetService) Create(c context.Context) (id string, err error) {
	serv.asset.Validate(c)
	return serv.db.CreateAsset(c, serv.asset)
}
