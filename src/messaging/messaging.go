package messaging

type IMessaging interface {
	Init() error
	Send(topic string, event string, data any) error
}
