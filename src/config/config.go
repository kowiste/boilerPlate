package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Name       string `envconfig:"NAME" validate:"required"` //service name
	Port       string `envconfig:"PORT" validate:"required"`
	APITimeOut int    `envconfig:"API_TOUT" validate:"required"`
	//Database
	DBType       string `envconfig:"DB_TYPE" validate:"required"`
	DBMock       bool   `envconfig:"DB_MOCK"`
	DBConnection string `envconfig:"DB_CONNECTION" validate:"dbValidation,required_if=DBMock false "`

	//Broker
	BrokerAddress string            `envconfig:"BROKER_ADDR"`
	LogLevel      string            `envconfig:"LOG_LEVEL"`
	ConsumerTopic map[string]string `envconfig:"CONSUMER_TOPICS"` //Key: topic, Value:GroupID
	ResponseTopic string            `envconfig:"RESPONSE_TOPIC"`
	LogTopic      string            `envconfig:"LOG_TOPIC"`
	Topic         []string          `envconfig:"TOPIC"`
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
