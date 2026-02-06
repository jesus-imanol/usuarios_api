package application

import (
	"productos-api/src/users/domain/entities"
	"productos-api/src/users/domain/repositories"
)

type RegisterUserUseCase struct {
	db repositories.IUser
}

func NewRegisterUserUseCase(db repositories.IUser) *RegisterUserUseCase {
	return &RegisterUserUseCase{db: db}
}

func (ru *RegisterUserUseCase) Execute(fullName, email, passwordHash string) (*entities.User, error) {
	user := entities.NewUser(fullName, email, passwordHash)
	err := ru.db.Register(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
