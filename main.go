package main

import (
	"log"
	dependenciesproduct "productos-api/src/products/infraestructure/dependencies_product"
	dependenciesuser "productos-api/src/users/infraestructure/dependencies_user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	dependenciesproduct.InitProduct(r)
	dependenciesuser.InitUsers(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
