package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kowiste/boilerplate/src/model/asset"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMockDB creates a new mock database connection and returns it along with the SQLMock instance.
func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("failed to open sqlmock database connection")
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to initialize gorm with sqlmock")
	}

	return gormDB, mock
}
func TestCreateAsset(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	asset := &asset.Asset{
		ID:          "1",
		ParentID:    "0",
		Description: "Test Description",
	}

	// Set up the expected behavior of the mock.
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `assets`").WithArgs(asset.ID, asset.ParentID, asset.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	id, err := repo.CreateAsset(ctx, asset)

	// Use the assert package to perform the test assertions.
	assert.NoError(t, err)
	assert.Equal(t, "1", id)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestAssets(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "description"}).
		AddRow("1", "0", "Test Description")

	mock.ExpectQuery("^SELECT \\* FROM `assets`").WillReturnRows(rows)

	assets, err := repo.Assets(ctx)

	assert.NoError(t, err)
	assert.Len(t, assets, 1)
	assert.Equal(t, "1", assets[0].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestAssetByID(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "description"}).
		AddRow("1", "0", "Test Description")

	mock.ExpectQuery("^SELECT \\* FROM `assets` WHERE id = \\? ORDER BY `assets`.`id` LIMIT \\?$").
		WithArgs("1", 1).
		WillReturnRows(rows)

	asset, err := repo.AssetByID(ctx, "1")

	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.Equal(t, "1", asset.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestUpdateAsset(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	asset := &asset.Asset{
		ID:          "1",
		ParentID:    "0",
		Description: "Updated Description",
	}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `assets` SET `parent_id`=\\?,`description`=\\? WHERE id = \\? AND `id` = \\?$").
		WithArgs(asset.ParentID, asset.Description, asset.ID, asset.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.UpdateAsset(ctx, asset)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestDeleteAsset(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `assets` WHERE id = \\?").WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteAsset(ctx, "1")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
