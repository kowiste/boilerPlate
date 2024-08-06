package nats

import (
	"errors"
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

	return nil
}
func (n *Nats) Consume(consumeTopics []string) error {
	if n.messageEvent == nil {
		return errors.New("should set message")
	}
	for _, topic := range consumeTopics {
		go n.consume(topic)
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

	for i := range n.config.Topic {
		_, err := n.js.AddStream(&nats.StreamConfig{
			Name:   n.config.Topic[i],
			MaxAge: time.Hour,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
