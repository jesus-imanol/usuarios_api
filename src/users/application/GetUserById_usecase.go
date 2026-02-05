package application

import (
	"productos-api/src/users/domain/entities"
	"productos-api/src/users/domain/repositories"
)

type GetUserByIdUseCase struct {
	db repositories.IUser
}

func NewGetUserById(db repositories.IUser) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{db: db}
}

func (gub *GetUserByIdUseCase) Execute(id int32) (*entities.User, error) {
	user, err := gub.db.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
