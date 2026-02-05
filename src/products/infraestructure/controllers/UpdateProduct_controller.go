package controllers

import (
	"net/http"
	"productos-api/src/products/application"
	"productos-api/src/products/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	updateProductUseCase *application.UpdateProductUseCase
}

func NewUpdateProductController(updateUseCase *application.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{
		updateProductUseCase: updateUseCase,
	}
}

func (upc *UpdateProductController) UpdateProduct(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var product *entities.Product
	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := upc.updateProductUseCase.Execute(int32(id), product.Name, product.Price, product.Quantity)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": gin.H{
			"type": "products",
			"id":   updatedProduct.Id,
			"attributes": gin.H{
				"name":     updatedProduct.Name,
				"price":    updatedProduct.Price,
				"quantity": updatedProduct.Quantity,
			},
		},
	}
	g.JSON(http.StatusOK, response)
}
