package repository

import (
	"boiler/src/model/mocks"
	"boiler/src/model/user"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()

	mockUsers := []user.User{
		{ID: "1", Name: "John Doe"},
		{ID: "2", Name: "Jane Smith"},
	}

	// Setup expectations
	mockRepo.On("Users", ctx).Return(mockUsers, nil)
	_, err := Get()
	assert.NotNil(t, err, "Singleton should return error")

	New(mockRepo)
	_, err = Get()
	assert.Nil(t, err)

}
