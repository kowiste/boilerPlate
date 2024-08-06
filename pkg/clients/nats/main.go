package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kowiste/boilerplate/pkg/config"
	"github.com/kowiste/boilerplate/pkg/nats"
	"github.com/kowiste/boilerplate/pkg/stream"
	"github.com/kowiste/boilerplate/src/model/user"
)

func main() {
	n := nats.New()
	err := n.Init(&config.ConfigBroker{
		Producer: "example",
		Address:  "nats://localhost:4222",
		Topic:    []string{"testNats"},
	})
	if err != nil {
		panic(err)
	}
	n.SetMessageEvent(onMessage)
	err = n.Consume([]string{"testNats"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("=  waiting for message = ")

	for {
		time.Sleep(time.Second)
	}
}

func onMessage(data []byte) {
	msg := new(stream.Message)
	json.Unmarshal(data, msg)

	fmt.Printf("- message - : %s\n", msg)
	u := new(user.User)
	json.Unmarshal(msg.Data, &u)
	fmt.Printf("- user updated - : %+v\n", u)
}
