package mongo

import (
	"context"

	"github.com/kowiste/boilerplatesrc/model/asset"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateAsset(a *asset.Asset) (id string, err error) {
	collection := m.db.Collection(a.TableName())
	_, err = collection.InsertOne(context.Background(), a)
	if err != nil {
		return "", err
	}
	id = a.ID
	return
}

func (m *MongoDB) GetAssets() (assets asset.Assets, err error) {

	collection := m.db.Collection(asset.Asset{}.TableName())
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var a asset.Asset
		err := cursor.Decode(&a)
		if err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return
}

func (m *MongoDB) GetAssetByID(id string) (a *asset.Asset, err error) {
	collection := m.db.Collection(a.TableName())
	err = collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&a)
	if err != nil {
		return nil, err
	}
	return
}

func (m *MongoDB) UpdateAsset(a *asset.Asset) (err error) {
	collection := m.db.Collection(a.TableName())
	_, err = collection.UpdateOne(context.Background(), bson.M{"id": a.ID}, bson.M{"$set": a})
	if err != nil {
		return err
	}
	return
}

func (m *MongoDB) DeleteAsset(id string) (err error) {
	collection := m.db.Collection(asset.Asset{}.TableName())
	_, err = collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return
}
