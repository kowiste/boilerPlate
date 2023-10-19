package nats

import (
	"encoding/json"
	"serviceX/src/config"
	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/nats-io/nats.go"
)

func (n *broker) Consume(topic string) {
	sub, err := n.js.Subscribe(topic, n.consumeEvent, nats.Durable(n.producer), nats.ManualAck())
	if err != nil {
		sub.Unsubscribe()
		sub.Drain()
		n.Reconnect(topic)
	}
}

func (n *broker) consumeEvent(msg *nats.Msg) {
	go func() {
		if n.messageEvent != nil {
			if n.messageEvent(msg.Data) != nil {
				msg.Nak() //error send again the message
			} else {
				msg.Ack()
			}
		}
	}()
}

// WriteLog write the log in the log topic
func (n *broker) WriteLog(data *log.LogEntry) {
	n.js.Publish(config.Get().LogTopic, data.Marshal())
}

// WriteMessage write a message in a specific topic
func (n *broker) WriteMessage(topic string, msg *model.Message) (err error) {
	msg.Producer = n.producer
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}
	_, err = n.js.Publish(topic, data)
	return
}

func (n *broker) Close() (err error) {
	err = n.conn.Drain()
	if err != nil {
		return
	}
	n.conn.Close()
	return
}
