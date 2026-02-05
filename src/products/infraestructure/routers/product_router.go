package routers

import (
	"productos-api/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRouter(
	r *gin.Engine,
	createProductController *controllers.CreateProductController,
	updateProductController *controllers.UpdateProductController,
	getProductsController *controllers.GetProductsController,
	getProductByIdController *controllers.GetProductByIdController,
	deleteProductController *controllers.DeleteProductController,
) {
	v1 := r.Group("/v1/products")
	{
		v1.POST("/", createProductController.CreateProduct)
		v1.PUT("/:id", updateProductController.UpdateProduct)
		v1.GET("/", getProductsController.GetAllProducts)
		v1.GET("/:id", getProductByIdController.GetProductById)
		v1.DELETE("/:id", deleteProductController.DeleteProduct)
	}
}
