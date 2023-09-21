package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	type dert struct {
		Field1 int `json:"field1" validate:"gt=3"`
	}
	p := dert{
		Field1: 0,
	}
	err := New(p)
	assert.Error(t, err)
	p.Field1=4
	err = New(p)
	assert.Nil(t, err)
}
