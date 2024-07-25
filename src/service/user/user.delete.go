package userservice

import (
	"context"
)

// Delete removes a user from the database by ID.
// Parameters:
// - c: The context for the operation.
// - id: The user ID to delete.
// Returns:
// - err: An error if the deletion fails.
func (serv UserService) Delete(c context.Context, id string) (err error) {
	return serv.db.DeleteUser(c, id)
}
