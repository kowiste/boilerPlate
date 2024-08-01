package nats

import (
	"sync"
	"time"

	"github.com/kowiste/boilerplate/pkg/config"
	"github.com/nats-io/nats.go"
)

type Nats struct {
	conn          *nats.Conn
	js            nats.JetStreamContext
	reconnCounter byte
	config        *config.ConfigBroker
	messageEvent  func(msg []byte)
	errorEvent    func(errorLog error)
	wg            sync.WaitGroup
}

func New() *Nats {
	return &Nats{}
}
func (n *Nats) Init(cfg *config.ConfigBroker) error {

	n.config = cfg
	err := n.connectCore()
	if err != nil {
		return err
	}
	n.createStreams()
	for _, topic := range n.config.ConsumerTopic {
		go n.Consume(topic)
	}
	return nil
}

func (n *Nats) connectCore() error {
	var err error
	n.wg.Add(1)
	n.conn, err = nats.Connect(n.config.Address, nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
		n.errorEvent(err)
	}), nats.ClosedHandler(func(_ *nats.Conn) {
		n.wg.Done()
	}))
	if err != nil {
		return err
	}
	n.js, err = n.conn.JetStream()
	if err != nil {
		return err
	}
	return nil
}

func (n *Nats) SetMessageEvent(callback func(msg []byte)) {
	n.messageEvent = callback
}
func (n *Nats) SetErrorEvent(callback func(errorLog error)) {
	n.errorEvent = callback
}
func (n *Nats) createStreams() (err error) {

	if n.config.ResponseTopic != "" {
		_, err = n.js.AddStream(&nats.StreamConfig{
			Name:   n.config.ResponseTopic,
			MaxAge: time.Hour,
		})
		if err != nil {
			return err
		}
	}

	for i := range n.config.Topic {
		_, err := n.js.AddStream(&nats.StreamConfig{
			Name:   n.config.Topic[i],
			MaxAge: time.Hour,
		})
		if err != nil {
			return err
		}
	}
	for _, topic := range n.config.ConsumerTopic {
		_, err := n.js.AddStream(&nats.StreamConfig{
			Name:   topic,
			MaxAge: time.Hour,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
