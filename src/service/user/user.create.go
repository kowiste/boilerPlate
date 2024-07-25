package userservice

import (
	"context"
)

// Create adds a new user to the database.
// Parameters:
// - c: The context for the operation.
// Returns:
// - id: The newly created user's ID.
// - err: An error if the creation fails.
func (serv UserService) Create(c context.Context) (id string, err error) {
	err = serv.user.Validate(c)
	if err != nil {
		return
	}
	return serv.db.CreateUser(c, serv.user)
}
