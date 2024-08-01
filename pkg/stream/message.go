package stream

type Message struct {
	ID       string
	UserID   string
	Producer string
	Event    string
	Data     any
}
