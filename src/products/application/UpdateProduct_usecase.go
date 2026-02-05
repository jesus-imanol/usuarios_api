package application

import (
	"productos-api/src/products/domain/entities"
	"productos-api/src/products/domain/repositories"
)

type UpdateProductUseCase struct {
	db repositories.IProductRepository
}

func NewUpdateProductUseCase(db repositories.IProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{db: db}
}

func (up *UpdateProductUseCase) Execute(id int32, name string, price float64, quantity int32) (*entities.Product, error) {
	product := &entities.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	err := up.db.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
