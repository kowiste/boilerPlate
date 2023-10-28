package nosql

import (
	"context"
	"net/http"
	"reflect"

	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s db) Create(c *gin.Context, data model.ModelI) {
	name := reflect.TypeOf(data).Elem().Name() // using the struct name as a collection name
	data.OnCreate()
	_, err := s.conn.Collection(name).InsertOne(context.Background(), data, nil)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data)
}

// FindOne
func (s db) FindOne(c *gin.Context, data model.ModelI) {
	name := reflect.TypeOf(data).Elem().Name() // using the struct name as a collection name
	oID, err := primitive.ObjectIDFromHex(c.Param("id	"))
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	filter := bson.D{{Key: "_id", Value: oID}}
	res := s.conn.Collection(name).FindOne(context.Background(), filter, nil)
	if res.Err() != nil {
		log.Get().Print(log.ErrorLevel, res.Err().Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	err = res.Decode(&data)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data)
}

// FindAll
func (s db) FindAll(c *gin.Context, request model.FindAllRequest, modelType model.ModelI, data any) {
	//name := reflect.TypeOf(data).Elem().Name() // using the struct name as a collection name

}
func (s db) Update(c *gin.Context, modelType model.ModelI, data map[string]any) {
	//name := reflect.TypeOf(data).Elem().Name() // using the struct name as a collection name
	modelType.OnUpdate()
}
func (s db) Delete(c *gin.Context, data model.ModelI) {
	//name := reflect.TypeOf(data).Elem().Name() // using the struct name as a collection name
	data.OnDelete()
}
