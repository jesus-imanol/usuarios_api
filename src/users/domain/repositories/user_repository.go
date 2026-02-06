package repositories

import "productos-api/src/users/domain/entities"

type IUser interface {
	Register(user *entities.User) error
	Update(id int32, fullname string, email string, passwordHash string) error
	GetAll() ([]*entities.User, error)
	Delete(id int32) error
	GetById(id int32) (*entities.User, error)
	UploadPicture(id int32, urlPicture string) error
	Login(email string) (*entities.User, error)
}
