package application

import "productos-api/src/users/domain/repositories"

type UpdateUserUseCase struct {
	db repositories.IUser
}

func NewUpdateUserUseCase(db repositories.IUser) *UpdateUserUseCase {
	return &UpdateUserUseCase{db: db}
}

func (uuc *UpdateUserUseCase) Execute(id int32, fullname string, email string, passwordHash string) error {
	err := uuc.db.Update(id, fullname, email, passwordHash)
	if err != nil {
		return err
	}
	return nil
}
