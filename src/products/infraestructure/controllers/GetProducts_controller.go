package controllers

import (
	"net/http"
	"productos-api/src/products/application"

	"github.com/gin-gonic/gin"
)

type GetProductsController struct {
	getProductsUseCase *application.GetProductsUseCase
}

func NewGetProductsController(getUseCase *application.GetProductsUseCase) *GetProductsController {
	return &GetProductsController{
		getProductsUseCase: getUseCase,
	}
}

func (gpc *GetProductsController) GetAllProducts(g *gin.Context) {
	products, err := gpc.getProductsUseCase.Execute()

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var productsData []gin.H
	for _, product := range products {
		productsData = append(productsData, gin.H{
			"type": "products",
			"id":   product.Id,
			"attributes": gin.H{
				"name":     product.Name,
				"price":    product.Price,
				"quantity": product.Quantity,
			},
		})
	}

	response := gin.H{
		"data": productsData,
	}
	g.JSON(http.StatusOK, response)
}
