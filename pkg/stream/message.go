package stream

import "encoding/json"

type Message struct {
	ID       string
	UserID   string
	Producer string
	Event    string
	Data     json.RawMessage
}
