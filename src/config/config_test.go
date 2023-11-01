package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := CreateInstance()
	assert.NotNil(t, err)
}

func TestGood(t *testing.T) {
	//Create .env
	f, _ := os.Create(".env")
	defer os.Remove(".env")
	f.WriteString("MY_ENV_VAR=my_env_value")
	f.Close()

	err := CreateInstance()
	assert.Nil(t, err)
}
