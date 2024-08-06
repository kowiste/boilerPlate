package config

type ConfigBroker struct {
	Producer string   `json:"BrokerProducer" env:"BROKER_PRODUCER"`
	Address  string   `json:"BrokerAddress" env:"BROKER_ADDRESS"`
	Topic    []string `json:"BrokerTopic" env:"BROKER_TOPIC"` //Topic that want to create
}
