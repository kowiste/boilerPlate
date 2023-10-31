package sql

import (
	"strings"
	"sync"

	"serviceX/src/config"
	"serviceX/src/handler/log"
	"serviceX/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type db struct {
	conn *gorm.DB
}

var lock = &sync.Mutex{}
var singleInstance *db

func CreatePostgres(dst ...interface{}) *db {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &db{}
		var err error
		if !config.Get().DBMock {
			// Establish the database connection
			singleInstance.conn, err = gorm.Open(postgres.Open(config.Get().DBSQL), &gorm.Config{
				//Logger:         logdb.Default.LogMode(logdb.Info),
				NamingStrategy: NamingStrategy{},
			})
			if err != nil {
				log.Get().Print(log.ErrorLevel, err.Error())
				panic(err)
			}

			// Automigrate the struct that are pass
			err = singleInstance.conn.AutoMigrate(dst...)
			if err != nil {
				panic(err)
			}
		}

	}
	return singleInstance
}

func Get() *db {
	return singleInstance
}

type NamingStrategy struct {
	schema.NamingStrategy
}

// ColumnName Modify the name of the columns
func (ns NamingStrategy) ColumnName(table, column string) string {
	if column == "" {
		return ""
	}
	if ns.NameReplacer != nil {
		tmpName := ns.NameReplacer.Replace(column)
		if tmpName == "" {
			return column
		}
		column = tmpName
	}
	if strings.ToLower(column) == "id" { //No matter how id is write, column will be id
		return "id"
	}
	return strings.ToLower(column[:1]) + column[1:] //First character allways a lower case
}

func (s db) validSchema(body map[string]interface{}, desc model.ModelI) bool {
	sch, err := schema.Parse(desc, &sync.Map{}, NamingStrategy{})
	if err != nil {
	}
	for key := range body {
		// TODO: review uploadFiles logic
		if sch.LookUpField(key) == nil && key != "uploadFiles" {
			return false
		}
	}

	return true
}

func (s db) Close() {
	db, _ := s.conn.DB()
	db.Close()
}
