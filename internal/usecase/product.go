package usecase

import (
	"abc/internal/model"
	"abc/internal/repository"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (u *ProductUsecase) CreateProduct(product *model.Product) error {
	return u.repo.CreateProduct(product)
}

func (u *ProductUsecase) GetProducts() ([]*model.Product, error) {
	return u.repo.GetAllProducts()
}
