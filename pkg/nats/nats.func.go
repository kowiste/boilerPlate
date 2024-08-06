package nats

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/kowiste/boilerplate/pkg/stream"
	"github.com/nats-io/nats.go"
)

func (n *Nats) consume(topic string) {
	sub, err := n.js.Subscribe(topic, func(msg *nats.Msg) {
		go n.messageEvent(msg.Data)
	}, nats.Durable(n.config.Producer))
	if err != nil {
		sub.Unsubscribe()
		sub.Drain()
		n.Reconnect(topic)
	}
}
func (n *Nats) Reconnect(topic string) {
	if n.reconnCounter > 3 {
		n.conn.Drain()
		n.conn.Close()
		n.connectCore()
		n.consume(topic)
		n.reconnCounter++
		time.Sleep(time.Duration(n.reconnCounter))
	}
}

func (n *Nats) Write(topic string, data []byte) (err error) {
	_, err = n.js.Publish(topic, data)
	if err != nil {
		return err
	}

	return nil
}

func (n Nats) WriteMessage(ID, userID, topic, event string, data any) (err error) {

	if n.config.Producer == "" || event == "" || data == nil {
		return errors.New("missing")
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	msg := stream.Message{
		ID:       ID,
		UserID:   userID,
		Producer: n.config.Producer,
		Event:    event,
		Data:     json.RawMessage(dataBytes),
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}
	err = n.Write(topic, b)
	if err != nil {
		return
	}
	return nil
}

func (n *Nats) Close() (err error) {
	err = n.conn.Drain()
	if err != nil {
		return
	}
	n.wg.Wait()
	n.conn.Close()
	return nil
}
