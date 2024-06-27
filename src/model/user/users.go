package user

import (
	"boiler/src/db"
	"context"
)

type UsersGetFilter struct {
	//filters for get user list
}

func (u Users) Get(c context.Context, filter *UsersGetFilter) (err error) {
	database, err := db.Get()
	u, err = database.GetUsers()
	return

}
