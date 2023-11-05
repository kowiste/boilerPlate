package nosql

import (
	"context"
	"net/http"

	"serviceX/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s db) Create(data model.ModelI) (err error) {

	data.OnCreate()
	_, err = s.conn.Collection(data.GetName()).InsertOne(context.Background(), data, nil)
	return
}

// FindOne
func (s db) FindOne(data model.ModelI) (status int, err error) {
	oID, _ := primitive.ObjectIDFromHex(data.GetID())
	filter := bson.D{{Key: "_id", Value: oID}}
	res := s.conn.Collection(data.GetName()).FindOne(context.Background(), filter, nil)
	if res.Err() == mongo.ErrNoDocuments {
		return http.StatusBadRequest, model.ErrNoDocuments
	} else if res.Err() != nil {
		return http.StatusInternalServerError, res.Err()
	}

	err = res.Decode(data)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

// FindAll
func (s db) FindAll(filter map[string]string, request model.FindAllRequest, modelType model.ModelI, data any) (status int, count int64, err error) {
	//limit the request
	opts := options.Find().SetSkip(int64(request.Offset))
	opts.SetLimit(int64(request.Limit))

	res, err := s.conn.Collection(modelType.GetName()).Find(context.Background(), filter, opts)
	if err != nil {
		return http.StatusInternalServerError, 0, err
	}
	count = int64(res.RemainingBatchLength())
	err = res.All(context.Background(), data)
	if err != nil {
		return http.StatusBadRequest, 0, err
	}
	return http.StatusOK, count, nil
}

// Update
func (s db) Update(modelType model.ModelI, data map[string]any) (status int, err error) {
	modelType.OnUpdate()
	result, err := s.conn.Collection(modelType.GetName()).UpdateOne(context.Background(),
		bson.M{"_id": modelType.GetID()},
		bson.D{{Key: "$set", Value: data}})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if result.ModifiedCount == 0 {
		return http.StatusBadRequest, model.ErrNoDocuments
	}
	return http.StatusOK, nil
}
func (s db) Delete(data model.ModelI) (status int, err error) {
	data.OnDelete()
	filter := bson.M{"_id": data.GetID()}
	result, err := s.conn.Collection(data.GetName()).DeleteOne(context.Background(), filter)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if result.DeletedCount == 0 {
		return http.StatusBadRequest, model.ErrNoDocuments
	}
	return http.StatusOK, nil
}
