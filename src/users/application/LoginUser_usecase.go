package application

import (
	"errors"
	"os"
	"productos-api/src/users/domain/entities"
	"productos-api/src/users/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	db repositories.IUser
}

func NewLoginUserUseCase(db repositories.IUser) *LoginUserUseCase {
	return &LoginUserUseCase{db: db}
}

func (lu *LoginUserUseCase) Execute(email string, passwordHash string) (*entities.User, error) {
	user, err := lu.db.Login(email)
	if err != nil {
		return nil, err
	}

	secretKey := os.Getenv("SECRET_KEY")
	passwordWithKey := passwordHash + secretKey

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(passwordWithKey))
	if err != nil {
		return nil, errors.New("credenciales invalidas")
	}

	return user, nil
}
