package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	passwordWithKey := password + secretKey

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordWithKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
