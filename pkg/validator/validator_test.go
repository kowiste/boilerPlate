package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {

	_, err := Get()
	assert.NotNil(t, err)
	v1 := New()
	v2, err := Get()
	assert.Nil(t, err)
	assert.Equal(t, v1, v2)
}
