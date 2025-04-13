package repository

import "abc/internal/model"

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetAllProducts() ([]*model.Product, error)
}
