package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	userservice "github.com/kowiste/boilerplate/src/service/user"
)

type Kafka struct {
	service *userservice.UserService
}

func New() (api *Kafka) {
	return &Kafka{}
}

func (k Kafka) Init() (err error) {
	//init kafka bla bla
	return
}

//	{
//		event:""
//		payload: any
//		time: ""
//	}
func (k Kafka) onMessage(data []byte) {
	dataEvent := ""
	switch dataEvent {
	case "createUser":
		k.createUser(data)
	case "updateUser":
		k.updateUser(data)
	default:
		fmt.Println("event not implemented", dataEvent)
	}
}
func (k Kafka) createUser(data []byte) {
	//modify data from kafka to match user create model
	user := k.service.GetUser()
	json.Unmarshal(data, user)
	ctx := context.Background()
	id, err := k.service.Create(ctx)
	if err != nil {
		//some error
	}
	fmt.Println("created with id", id)
}
func (k Kafka) updateUser(data any) {

}
