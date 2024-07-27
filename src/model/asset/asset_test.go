package asset

import (
	"context"
	"testing"

	"github.com/kowiste/boilerplate/pkg/validator"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	asset := new(Asset)
	c := context.Background()

	assert.Equal(t, "assets", asset.TableName(), "asset table should be assets")
	//validator not set
	err := asset.Validate(c)
	assert.Equal(t, "validator not set", err.Error())
	//validation error
	validator.New()
	err = asset.Validate(c)
	assert.NotNil(t, err)
	//Ok
	asset.ParentID = "d3e8d6b0-1234-4f8a-93d6-09a8f0dc2a7e"
	err = asset.Validate(c)
	assert.Nil(t, err)
}
func TestAssetOther(t *testing.T) {
	user := Asset{
		ID:          "345",
		ParentID:    "34",
		Description: "test description",
	}
	uPB := user.ToGRPC()
	assert.Equal(t, "345", uPB.Id)
	assert.Equal(t, "34", uPB.ParentId)
	assert.Equal(t, "test description", uPB.Description)

}

func TestAssetsOther(t *testing.T) {
	users := new(Assets)

	uu := []Asset{{
		ID:          "345",
		ParentID:    "34",
		Description: "test description",
	}}
	users = (*Assets)(&uu)
	uPB := users.ToGRPC()
	assert.Equal(t, "345", uPB[0].Id)
	assert.Equal(t, "34", uPB[0].ParentId)
	assert.Equal(t, "test description", uPB[0].Description)

}
