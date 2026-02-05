package application

import (
	"productos-api/src/products/domain/entities"
	"productos-api/src/products/domain/repositories"
)

type CreateProductUseCase struct {
	db repositories.IProductRepository
}

func NewCreateProductUseCase(db repositories.IProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{db: db}
}

func (cp *CreateProductUseCase) Execute(name string, price float64, quantity int32) (*entities.Product, error) {
	product := entities.NewProduct(name, price, quantity)
	err := cp.db.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
