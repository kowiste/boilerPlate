package sql

import (
	"strings"
	"sync"

	"serviceX/src/config"
	"serviceX/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
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
			singleInstance.conn, err = gorm.Open(postgres.Open(config.Get().DBConnection), &gorm.Config{
				Logger:         log.Default.LogMode(log.Info),
				NamingStrategy: NamingStrategy{},
			})
			if err != nil {
				panic(err)
			}
			err = singleInstance.conn.
				AutoMigrate(dst...) //Automigrate the struct that are pass
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

func (s db) validSchema(body map[string]interface{}, desc model.ModelInterface) bool {
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