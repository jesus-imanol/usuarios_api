package application

import (
	"productos-api/src/products/domain/repositories"
)

type DeleteProductUseCase struct {
	db repositories.IProductRepository
}

func NewDeleteProductUseCase(db repositories.IProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{db: db}
}

func (dp *DeleteProductUseCase) Execute(id int32) error {
	err := dp.db.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
