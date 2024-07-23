package mysql

import (
	"boiler/src/config"
	"boiler/src/model/asset"
	"boiler/src/model/user"
	"boiler/src/mysql"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mock configuration
type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) Get[T any]() (T, error) {
	args := m.Called()
	return args.Get(0).(T), args.Error(1)
}

// Mock GORM DB
type MockDB struct {
	mock.Mock
}

func (m *MockDB) AutoMigrate(dst ...interface{}) error {
	args := m.Called(dst...)
	return args.Error(0)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Model(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	args := m.Called(query, args)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Updates(values interface{}) *gorm.DB {
	args := m.Called(values)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds...)
	return args.Get(0).(*gorm.DB)
}

func TestInit(t *testing.T) {
	mockConfig := new(MockConfig)
	mockDB := new(MockDB)
	cfg := config.BoilerConfig{
		DatabaseUser:     "user",
		DatabasePassword: "password",
		DatabaseURL:      "localhost",
		DatabasePort:     "3306",
		DatabaseName:     "test_db",
	}
	mockConfig.On("Get").Return(cfg, nil)

	mysqlConn := &mysql.MySQL{
		db: &gorm.DB{},
	}
	err := mysqlConn.Init()
	assert.NoError(t, err)

	mockConfig.AssertExpectations(t)
	mockDB.AssertExpectations(t)
}

func TestCreateAsset(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: nil})

	asset := &asset.Asset{
		ID:   "1",
		Name: "Test Asset",
	}
	id, err := mysqlConn.CreateAsset(context.Background(), asset)
	assert.NoError(t, err)
	assert.Equal(t, "1", id)

	mockDB.AssertExpectations(t)
}

func TestAssets(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Find", mock.Anything).Return(&gorm.DB{Error: nil})

	assets, err := mysqlConn.Assets(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, assets)

	mockDB.AssertExpectations(t)
}

func TestAssetByID(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{Error: nil})

	asset, err := mysqlConn.AssetByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.NotNil(t, asset)

	mockDB.AssertExpectations(t)
}

func TestUpdateAsset(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Model", mock.Anything).Return(mockDB)
	mockDB.On("Where", mock.Anything, mock.Anything).Return(mockDB)
	mockDB.On("Updates", mock.Anything).Return(&gorm.DB{Error: nil})

	asset := &asset.Asset{
		ID:   "1",
		Name: "Updated Asset",
	}
	err := mysqlConn.UpdateAsset(context.Background(), asset)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

func TestDeleteAsset(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{Error: nil})

	err := mysqlConn.DeleteAsset(context.Background(), "1")
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: nil})

	user := &user.User{
		ID:   "1",
		Name: "Test User",
	}
	id, err := mysqlConn.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, "1", id)

	mockDB.AssertExpectations(t)
}

func TestUsers(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Find", mock.Anything).Return(&gorm.DB{Error: nil})

	users, err := mysqlConn.Users(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, users)

	mockDB.AssertExpectations(t)
}

func TestUserByID(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{Error: nil})

	user, err := mysqlConn.UserByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	mockDB.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Model", mock.Anything).Return(mockDB)
	mockDB.On("Where", mock.Anything, mock.Anything).Return(mockDB)
	mockDB.On("Updates", mock.Anything).Return(&gorm.DB{Error: nil})

	user := &user.User{
		ID:   "1",
		Name: "Updated User",
	}
	err := mysqlConn.UpdateUser(context.Background(), user)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockDB := new(MockDB)
	mysqlConn := &mysql.MySQL{
		db: mockDB,
	}
	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{Error: nil})

	err := mysqlConn.DeleteUser(context.Background(), "1")
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}
