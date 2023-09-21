package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"test.com/validator"
)

type Config struct {
	Port             string `envconfig:"PORT" validate:"required"`
	DBGORMConnection string `envconfig:"DB_GORM_CONNECTION" validate:"required"`
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
		return validator.New(singleInstance)
	}
	return nil
}

func GetInstance() *Config {
	return singleInstance
}
