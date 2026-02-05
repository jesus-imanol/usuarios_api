package dependenciesproduct

import (
	"productos-api/src/products/application"
	"productos-api/src/products/infraestructure/adapters"
	"productos-api/src/products/infraestructure/controllers"
	"productos-api/src/products/infraestructure/routers"

	"github.com/gin-gonic/gin"
)

func InitProduct(r *gin.Engine) {
	// Inicializar adaptador MySQL
	mysql, err := adapters.NewMySQL()
	if err != nil {
		panic(err)
	}

	// Inicializar casos de uso
	createProductUseCase := application.NewCreateProductUseCase(mysql)
	updateProductUseCase := application.NewUpdateProductUseCase(mysql)
	getProductsUseCase := application.NewGetProductsUseCase(mysql)
	getProductByIdUseCase := application.NewGetProductByIdUseCase(mysql)
	deleteProductUseCase := application.NewDeleteProductUseCase(mysql)

	// Inicializar controladores
	createProductController := controllers.NewCreateProductController(createProductUseCase)
	updateProductController := controllers.NewUpdateProductController(updateProductUseCase)
	getProductsController := controllers.NewGetProductsController(getProductsUseCase)
	getProductByIdController := controllers.NewGetProductByIdController(getProductByIdUseCase)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)

	// Inicializar router
	routers.ProductRouter(
		r,
		createProductController,
		updateProductController,
		getProductsController,
		getProductByIdController,
		deleteProductController,
	)
}
