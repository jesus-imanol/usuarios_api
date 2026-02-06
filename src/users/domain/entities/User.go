package entities

type User struct {
	Id           int32  `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func NewUser(fullName, email, passwordHash string) *User {
	return &User{
		FullName:     fullName,
		Email:        email,
		PasswordHash: passwordHash,
	}
}
