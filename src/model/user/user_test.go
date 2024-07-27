package user

import (
	"context"
	"testing"

	"github.com/kowiste/boilerplate/pkg/validator"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	user := new(User)
	c := context.Background()

	assert.Equal(t, "users", user.TableName(), "user table should be users")
	//validator not set
	err := user.Validate(c)
	assert.Equal(t, "validator not set", err.Error())
	//validation error
	validator.New()
	err = user.Validate(c)
	assert.NotNil(t, err)
	//Ok
	user.Name = "Pablo"
	user.LastName = "Garcia"
	user.Age = 45
	err = user.Validate(c)
	assert.Nil(t, err)
}
func TestUserOther(t *testing.T) {
	user := User{
		ID:       "345",
		Name:     "Pablo",
		LastName: "Garcia",
		Age:      35,
	}
	uPB := user.ToGRPC()
	assert.Equal(t, "345", uPB.Id)
	assert.Equal(t, "Pablo", uPB.Name)
	assert.Equal(t, "Garcia", uPB.LastName)
	assert.Equal(t, uint32(35), uPB.Age)

}

func TestUsersOther(t *testing.T) {
	users := new(Users)

	uu := []User{{
		ID:       "345",
		Name:     "Pablo",
		LastName: "Garcia",
		Age:      35,
	}}
	users = (*Users)(&uu)
	uPB := users.ToGRPC()
	assert.Equal(t, "345", uPB[0].Id)
	assert.Equal(t, "Pablo", uPB[0].Name)
	assert.Equal(t, "Garcia", uPB[0].LastName)
	assert.Equal(t, uint32(35), uPB[0].Age)

}
