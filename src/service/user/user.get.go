package userservice

import (
	"context"

	"github.com/kowiste/boilerplate/src/model/user"
)

// Users retrieves a list of users from the database based on the input criteria.
// Parameters:
// - c: The context for the operation.
// - input: The criteria for finding users.
// Returns:
// - users: A list of users that match the criteria.
// - err: An error if the retrieval fails.
func (serv UserService) Users(c context.Context, input *user.FindUsersInput) (users user.Users, err error) {
	return serv.db.Users(c, input)
}

// UserByID retrieves a user from the database by their ID.
// Parameters:
// - c: The context for the operation.
// - id: The user ID to find.
// Returns:
// - users: The user with the specified ID.
// - err: An error if the retrieval fails.
func (serv UserService) UserByID(c context.Context, id string) (users *user.User, err error) {
	return serv.db.UserByID(c, id)
}
