package repositories

import "productos-api/src/products/domain/entities"

type IProductRepository interface {
	CreateProduct(product *entities.Product) error
	UpdateProduct(product *entities.Product) error
	GetProductById(id int32) (*entities.Product, error)
	GetAllProducts() ([]*entities.Product, error)
	DeleteProduct(id int32) error
}
