package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port         string `envconfig:"PORT" validate:"required"`
	DBType       string `envconfig:"DB_TYPE" validate:"required"`
	DBMock       bool   `envconfig:"DB_MOCK"`
	DBConnection string `envconfig:"DB_CONNECTION" validate:"dbValidation,required_if=DBMock false "`
}

var lock = &sync.Mutex{}
var singleInstance *Config

func CreateInstance(fileName ...string) error {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = new(Config)
		err := godotenv.Load(fileName...)
		if err != nil {
			return err
		}
		err = envconfig.Process("", singleInstance)
		if err != nil {
			return err
		}

		return nil
	}
	return nil
}

func Get() *Config {
	return singleInstance
}
