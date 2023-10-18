package nats

import (
	"fmt"
	"serviceX/src/config"
	"serviceX/src/handler/log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type broker struct {
	conn          *nats.Conn
	js            nats.JetStreamContext
	reconnCounter byte
	producer      string
	messageEvent  func(msg []byte) error
	errorEvent    func(errorLog error)
	wg            sync.WaitGroup
}

var lock = &sync.Mutex{}
var singleInstance *broker

func CreateInstance(producerName string) (err error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &broker{producer: producerName}
		err = singleInstance.connectCore()
		if err != nil {
			return
		}
		singleInstance.initializateStreams()
		for topic, group := range config.Get().ConsumerTopic {
			log.Get().Print(log.InfoLevel, fmt.Sprintf("Consuming %s in the group %s", topic, group))
			go singleInstance.Consume(topic)
		}
	}
	return
}

func Get() *broker {
	return singleInstance
}

func (n *broker) connectCore() (err error) {

	n.wg.Add(1)
	n.conn, err = nats.Connect(config.Get().BrokerAddress, nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
		n.errorEvent(err)
	}), nats.ClosedHandler(func(_ *nats.Conn) {
		n.wg.Done()
	}))
	if err != nil {
		return
	}
	n.js, err = n.conn.JetStream()
	if err != nil {
		return
	}
	return nil
}

func (n *broker) SetMessageEvent(callback func(msg []byte) error) {
	n.messageEvent = callback
}
func (n *broker) SetErrorEvent(callback func(errorLog error)) {
	n.errorEvent = callback
}

func (n *broker) Reconnect(topic string) {
	if n.reconnCounter > 3 {
		n.conn.Drain()
		n.conn.Close()
		n.connectCore()
		n.Consume(topic)
		n.reconnCounter++
		time.Sleep(time.Duration(n.reconnCounter))
	}
}

// initializateStreams create the streams that will be use for the app
func (n *broker) initializateStreams() (err error) {
	if config.Get().LogTopic != "" {
		_, err = n.js.AddStream(&nats.StreamConfig{
			Name:   config.Get().LogTopic,
			MaxAge: time.Hour,
		})
		if err != nil {
			return
		}
	}
	//Topic for response async request
	if config.Get().ResponseTopic != "" {
		_, err = n.js.AddStream(&nats.StreamConfig{
			Name:   config.Get().ResponseTopic,
			MaxAge: time.Hour,
		})
		if err != nil {
			return
		}
	}

	for i := range config.Get().Topic {
		_, err = n.js.AddStream(&nats.StreamConfig{
			Name:   config.Get().Topic[i],
			MaxAge: time.Hour,
		})
		if err != nil {
			return
		}
	}
	for topic := range config.Get().ConsumerTopic {
		_, err = n.js.AddStream(&nats.StreamConfig{
			Name:   topic,
			MaxAge: time.Hour,
		})
		if err != nil {
			return
		}
	}
	return nil
}
