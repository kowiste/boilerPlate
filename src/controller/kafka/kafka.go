package kafka

import (
	userservice "boiler/src/service/user"
	"context"
	"encoding/json"
	"fmt"
)

type Kafka struct {
	service *userservice.UserService
}

func New() (api *Kafka) {
	return &Kafka{}
}

func (k Kafka) Init() {
	//init kafka bla bla
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
	}
}
func (k Kafka) createUser(data []byte) {
	//modify data from kafka to match user create model
	user := k.service.GetUser()
	json.Unmarshal(data, user)
	ctx := context.Background()
	id, err := k.service.Create(ctx)
	if err != nil {
		//someting error
	}
	fmt.Println("created with id", id)
}
func (k Kafka) updateUser(data any) {

}
