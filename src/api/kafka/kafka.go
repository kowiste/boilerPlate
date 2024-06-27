package kafka

import (
	"boiler/src/model/user"
	"context"
	"encoding/json"
)

type Kafka struct{}

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
	user := new(user.User)
	json.Unmarshal(data,user)
	ctx := context.Background()
	err := user.Create(ctx)
	if err != nil {
		//someting error
	}
}
func (k Kafka) updateUser(data any) {

}
