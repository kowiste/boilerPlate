package mongo

import (
	"context"

	"github.com/kowiste/boilerplatesrc/model/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateUser(u *user.User) (id string, err error) {
	collection := m.db.Collection(u.TableName())
	_, err = collection.InsertOne(context.Background(), u)
	if err != nil {
		return "", err
	}
	id = u.ID
	return
}

func (m *MongoDB) GetUsers() (users []user.User, err error) {

	collection := m.db.Collection(user.User{}.TableName())
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var u user.User
		err := cursor.Decode(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return
}

func (m *MongoDB) GetUserByID(id string) (u *user.User, err error) {
	collection := m.db.Collection(u.TableName())
	err = collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return
}

func (m *MongoDB) UpdateUser(id string, u *user.User) (err error) {
	collection := m.db.Collection(u.TableName())
	_, err = collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": u})
	if err != nil {
		return err
	}
	return
}

func (m *MongoDB) DeleteUser(id string) (err error) {
	collection := m.db.Collection(user.User{}.TableName())
	_, err = collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return
}
