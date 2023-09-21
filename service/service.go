package service

import (
	"sync"

	"github.com/gin-gonic/gin"
	"test.com/handler/database/sql"
	"test.com/model"
)

type service struct {
	db Database
}

type Database interface {
	Create(*gin.Context, model.ModelInterface)
	FindOne(*gin.Context, model.ModelInterface)
	FindAll(*gin.Context, model.ModelInterface)
	Update(*gin.Context, model.ModelInterface)
	Delete(*gin.Context, model.ModelInterface)
}

var lock = &sync.Mutex{}

var singleInstance *service

func Init() *service {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &service{
				db: sql.CreateInstance(),
			}
		}
	}
	return singleInstance
}

func (s service) Authorization(c *gin.Context) {

}
