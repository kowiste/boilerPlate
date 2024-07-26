package asset

import (
	"context"
	"testing"

	"github.com/kowiste/boilerplatepkg/validator"

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
