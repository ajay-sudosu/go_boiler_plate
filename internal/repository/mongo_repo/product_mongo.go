package mongo_repo

import (
	"abc/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *MongoProductRepository) CreateProduct(product *model.Product) error {
	_, err := r.collection.InsertOne(context.TODO(), product)
	return err
}

func (r *MongoProductRepository) GetAllProducts() ([]*model.Product, error) {
	cur, err := r.collection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var products []*model.Product
	for cur.Next(context.TODO()) {
		var product model.Product
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
