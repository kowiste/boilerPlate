package nats

import (
	"github.com/google/uuid"
	"github.com/kowiste/boilerplate/pkg/nats"

	conf "github.com/kowiste/boilerplate/src/config"
	"github.com/kowiste/config"
)

type Nats struct {
	conn nats.Nats
}

func New() *Nats {
	return &Nats{
		conn: *nats.New(),
	}
}

func (n *Nats) Init() (err error) {
	c, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		return
	}
	return n.conn.Init(&c.ConfigBroker)
}
func (n Nats) Send(topic string, event string, data any) (err error) {
	return n.conn.WriteMessage(uuid.NewString(), "pablo", topic, event, data)
}
