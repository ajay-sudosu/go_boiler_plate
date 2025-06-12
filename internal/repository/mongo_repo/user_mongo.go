package mongo_repo

import (
	"abc/internal/model"
	"abc/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoNetworkRepository struct {
	collection *mongo.Collection
}

func NewNetworkRepository(db *mongo.Database) repository.NetworkRepository {
	return &MongoNetworkRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoNetworkRepository) CreateNetwork(user *model.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *MongoNetworkRepository) GetNetworks() ([]*model.User, error) {
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
