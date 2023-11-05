package main

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

const topic string = "command"

// message to send in nats
type send struct {
	Producer string `json:"producer"`
	Event    string `json:"event"`
	Model    string `json:"model"`
	Data     []byte `json:"data"`
}
type delete struct {
	ID int `json:"id"`
}

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := send{
		Producer: "test",
		Event:    "delete",
		Model:    "stuff",
	}
	s := delete{
		ID: 2,
	}
	msg.Data, _ = json.Marshal(s)
	b, _ := json.Marshal(msg)
	// Create a publisher
	err = nc.Publish(topic, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the connection
	nc.Close()
}
