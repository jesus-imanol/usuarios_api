package main

import (
	"log"
	"os" // Añadido para leer variables de entorno del sistema
	dependenciesproduct "productos-api/src/products/infraestructure/dependencies_product"
	dependenciesuser "productos-api/src/users/infraestructure/dependencies_user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Intentamos cargar el .env (útil para local), pero no matamos la app si falla
	// En Railway, godotenv fallará porque las variables ya están en el sistema
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: No se encontró archivo .env, usando variables de entorno del sistema")
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

	// OBTENER EL PUERTO DE RAILWAY
	// Railway asigna un puerto dinámico; si usamos :8080 fijo, la app no responderá
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto si estás corriendo en local
	}

	log.Printf("Servidor iniciando en el puerto %s", port)
	
	// Cambiamos ":8080" por la variable port
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Fallo al iniciar el servidor: ", err)
	}
}