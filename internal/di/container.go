package di

import (
	"abc/internal/handler"

	"abc/internal/repository/mongo_repo"
	"abc/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	NetworkHandler *handler.NetworkHandler
	ProductHandler *handler.ProductHandler
}

func NewContainer(db *mongo.Database) (*Container, error) {
	networkRepo := mongo_repo.NewNetworkRepository(db)
	// networkAdaptor := adaptor.NetworkAdapter()
	userUsecase := usecase.NewNetworkUsecase(networkRepo)
	networkHandler := handler.NewUserHandler(userUsecase)

	productRepo := mongo_repo.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	return &Container{
		NetworkHandler: networkHandler,
		ProductHandler: productHandler,
	}, nil
}
