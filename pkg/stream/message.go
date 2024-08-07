package stream

import "encoding/json"

type Message struct {
	ID       string
	UserID   string
	Producer string
	Event    string
	Data     []byte
}

func (m *Message) Marshal(data any) (out []byte, err error) {
	dataB, err := json.Marshal(data)
	if err != nil {
		return
	}
	m.Data = dataB
	return json.Marshal(m)
}

func (m *Message) Decode(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m Message) UnMarshal(target interface{}) error {
	return json.Unmarshal(m.Data, target)
}
