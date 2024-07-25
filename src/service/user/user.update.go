package userservice

import (
	"context"
)

// Update modifies an existing user in the database.
// Parameters:
// - c: The context for the operation.
// Returns:
// - err: An error if the update fails.
func (serv UserService) Update(c context.Context) (err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.UpdateUser(c, serv.user)
}
