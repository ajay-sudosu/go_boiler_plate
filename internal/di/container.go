package di

import (
	"abc/internal/handler"
	"abc/internal/repository/mongo_repo"
	"abc/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	UserHandler    *handler.UserHandler
	ProductHandler *handler.ProductHandler
}

func NewContainer(db *mongo.Database) (*Container, error) {
	userRepo := mongo_repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	productRepo := mongo_repo.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	return &Container{
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}, nil
}
