package user

import (
	"boiler/pkg/validator"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
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
	err = user.Validate(c)
	assert.Nil(t, err)
}
