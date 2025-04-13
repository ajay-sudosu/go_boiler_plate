package mongo_repo

import (
	"abc/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) CreateUser(user *model.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *MongoUserRepository) GetAllUsers() ([]*model.User, error) {
	cur, err := r.collection.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []*model.User
	for cur.Next(context.Background()) {
		var user model.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
