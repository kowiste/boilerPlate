package mongo

import (
	"context"
	"fmt"

	conf "boiler/src/config"

	"github.com/kowiste/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func New() *MongoDB {
	return &MongoDB{}
}

func (m *MongoDB) Init() (err error) {
	cnf, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cnf.DatabaseUser,
		cnf.DatabasePassword,
		cnf.DatabaseURL,
		cnf.DatabasePort,
	))

	m.client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = m.client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}

	m.db = m.client.Database(cnf.DatabaseName)

	return nil
}
