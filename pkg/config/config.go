package config

type ConfigBroker struct {
	Producer      string
	Address       string
	ResponseTopic string
	ConsumerTopic []string
	Topic         []string
}
