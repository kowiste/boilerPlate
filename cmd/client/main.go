package main

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

const topic string = "command"

// message to send in nats
type send struct {
}

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := send{}
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
