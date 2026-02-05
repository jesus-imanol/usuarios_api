package controllers

import (
	"net/http"
	"productos-api/src/products/application"
	"productos-api/src/products/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	createProductUseCase *application.CreateProductUseCase
}

func NewCreateProductController(createUseCase *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		createProductUseCase: createUseCase,
	}
}

func (cpc *CreateProductController) CreateProduct(g *gin.Context) {
	var product *entities.Product
	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := cpc.createProductUseCase.Execute(product.Name, product.Price, product.Quantity)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": gin.H{
			"type": "products",
			"id":   createdProduct.Id,
			"attributes": gin.H{
				"name":     createdProduct.Name,
				"price":    createdProduct.Price,
				"quantity": createdProduct.Quantity,
			},
		},
	}
	g.JSON(http.StatusCreated, response)
}
