package user

import (
	pbUser "github.com/kowiste/boilerplate/doc/proto/user"
)

type User struct {
	ID       string `json:"id" validate:"required,uuid"`
	Name     string `json:"name" validate:"gt=0,lte=130"`
	LastName string `json:"lastName" validate:"gt=0,lte=130"`
	Age      uint   `json:"age" validate:"gt=0"`
}
type Users []User

type FindUsersInput struct {
	Text string `json:"text"`
	Age  int    `json:"age"`
}

func (u User) TableName() string {
	return "users"
}
func (u User) ToGRPC() *pbUser.User {
	return &pbUser.User{
		Id:       u.ID,
		Name:     u.Name,
		LastName: u.LastName,
		Age:      uint32(u.Age),
	}
}

// ToGRPC converts the User slice to a User protobuf message.
func (u Users) ToGRPC() []*pbUser.User {
	assets := make([]*pbUser.User, len(u))
	for i, user := range u {
		assets[i] = user.ToGRPC()
	}

	return assets

}
