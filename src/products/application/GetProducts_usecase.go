package application

import (
	"productos-api/src/products/domain/entities"
	"productos-api/src/products/domain/repositories"
)

type GetProductsUseCase struct {
	db repositories.IProductRepository
}

func NewGetProductsUseCase(db repositories.IProductRepository) *GetProductsUseCase {
	return &GetProductsUseCase{db: db}
}

func (gp *GetProductsUseCase) Execute() ([]*entities.Product, error) {
	products, err := gp.db.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
