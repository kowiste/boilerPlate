package mocks

import (
	"context"

	"github.com/kowiste/boilerplatesrc/model/asset"
	"github.com/kowiste/boilerplatesrc/model/user"

	"github.com/stretchr/testify/mock"
)

// MockRepository struct implementing IRepository
type MockRepository struct {
	mock.Mock
}

// Init initializes the mock repository
func (m *MockRepository) Init() error {
	args := m.Called()
	return args.Error(0)
}

// CreateUser creates a new user
func (m *MockRepository) CreateUser(ctx context.Context, u *user.User) (string, error) {
	args := m.Called(ctx, u)
	return args.String(0), args.Error(1)
}

// Users returns all users
func (m *MockRepository) Users(ctx context.Context, input *user.FindUsersInput) (user.Users, error) {
	args := m.Called(ctx)
	if users, ok := args.Get(0).(user.Users); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

// UserByID returns a user by ID
func (m *MockRepository) UserByID(ctx context.Context, id string) (*user.User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*user.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateUser updates a user
func (m *MockRepository) UpdateUser(ctx context.Context, u *user.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

// DeleteUser deletes a user
func (m *MockRepository) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// CreateAsset creates a new asset
func (m *MockRepository) CreateAsset(ctx context.Context, a *asset.Asset) (string, error) {
	args := m.Called(ctx, a)
	return args.String(0), args.Error(1)
}

// Assets returns all assets
func (m *MockRepository) Assets(ctx context.Context) (asset.Assets, error) {
	args := m.Called(ctx)
	if assets, ok := args.Get(0).(asset.Assets); ok {
		return assets, args.Error(1)
	}
	return nil, args.Error(1)
}

// AssetByID returns an asset by ID
func (m *MockRepository) AssetByID(ctx context.Context, id string) (*asset.Asset, error) {
	args := m.Called(ctx, id)
	if asset, ok := args.Get(0).(*asset.Asset); ok {
		return asset, args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateAsset updates an asset
func (m *MockRepository) UpdateAsset(ctx context.Context, a *asset.Asset) error {
	args := m.Called(ctx, a)
	return args.Error(0)
}

// DeleteAsset deletes an asset
func (m *MockRepository) DeleteAsset(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
