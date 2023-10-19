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
	log           chan *log.LogEntry
}

var lock = &sync.Mutex{}
var singleInstance *broker

func CreateInstance(producerName string) (err error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &broker{producer: producerName, log: make(chan *log.LogEntry)}
		err = singleInstance.connectCore()
		if err != nil {
			return
		}
		singleInstance.initializateStreams()
		for topic, group := range config.Get().ConsumerTopic {
			log.Get().Print(log.InfoLevel, fmt.Sprintf("Consuming %s in the group %s", topic, group))
			go singleInstance.Consume(topic)
		}
		go singleInstance.run()
	}
	return
}

func Get() *broker {
	return singleInstance
}

// run waiting for log
func (n *broker) run() {
	for {
		l := <-n.log
		n.WriteLog(l)
	}
}

func (n *broker) connectCore() (err error) {
	n.conn, err = nats.Connect(config.Get().BrokerAddress)
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

// GetChannel return log channel for nats
func (n *broker) GetChannel() chan *log.LogEntry {
	return n.log
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
