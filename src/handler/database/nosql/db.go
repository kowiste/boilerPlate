package nosql

import (
	"context"
	"serviceX/src/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type db struct {
	conn *mongo.Database
}

var lock = &sync.Mutex{}
var singleInstance *db

func CreateMongo(databaseName string, dst ...interface{}) *db {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &db{}
		if !config.Get().DBMock {
			clientOptions := options.Client().ApplyURI(config.Get().DBConnection)
			conn, err := mongo.Connect(context.Background(), clientOptions)
			if err != nil {
				panic(err)
			}
			err = conn.Ping(context.Background(), readpref.Primary())
			if err != nil {
				panic(err)
			}
			singleInstance.conn = conn.Database(databaseName)
		}

	}
	return singleInstance
}
func Get() *db {
	return singleInstance
}

func (s db) Close() {
	s.conn.Client().Disconnect(context.Background())
}
