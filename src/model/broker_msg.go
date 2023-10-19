package model

import (
	"encoding/json"
	"serviceX/src/config"

	"github.com/google/uuid"
)

type BaseID struct {
	ID       string `json:"id"`
	TenantID string `json:"tenant_id"`
	UserID   string `json:"user_id"`
}
type Message struct {
	BaseID
	Producer string `json:"producer"`
	Event    string `json:"event"`
	Data     []byte `json:"data"`
}

func NewMessage(event string, data any) *Message {
	b, _ := json.Marshal(data)
	return &Message{
		BaseID: BaseID{
			ID: uuid.NewString(),
		},
		Producer: config.Get().Name,
		Event:    event,
		Data:     b,
	}
}
