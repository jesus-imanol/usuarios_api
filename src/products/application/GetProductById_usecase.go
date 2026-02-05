package application

import (
	"productos-api/src/products/domain/entities"
	"productos-api/src/products/domain/repositories"
)

type GetProductByIdUseCase struct {
	db repositories.IProductRepository
}

func NewGetProductByIdUseCase(db repositories.IProductRepository) *GetProductByIdUseCase {
	return &GetProductByIdUseCase{db: db}
}

func (gp *GetProductByIdUseCase) Execute(id int32) (*entities.Product, error) {
	product, err := gp.db.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
