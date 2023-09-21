package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	serv := Init()
	assert.NotNil(t, serv)
}
