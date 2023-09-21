package sql

import (
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"test.com/config"
)

type db struct {
	conn *gorm.DB
}

var lock = &sync.Mutex{}
var singleInstance *db

func CreateInstance(dst ...interface{}) *db {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		var err error
		singleInstance.conn, err = gorm.Open(mysql.Open(config.GetInstance().DBGORMConnection), &gorm.Config{
			Logger:         log.Default.LogMode(log.Info),
			NamingStrategy: NamingStrategy{},
		})
		if err != nil {
			panic(err)
		}
		err = singleInstance.conn.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_general_ci").AutoMigrate(dst...)
		if err != nil {
			panic(err)
		}
	}
	return singleInstance
}
func GetInstance() *db {
	return singleInstance
}

type NamingStrategy struct {
	schema.NamingStrategy
}

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
	return strings.ToLower(column[:1]) + column[1:]
}
