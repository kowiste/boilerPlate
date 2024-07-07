package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := Init()
	assert.NotNil(t, err, "The two words should be the same.")
}
