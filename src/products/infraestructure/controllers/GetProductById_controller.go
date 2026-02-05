package controllers

import (
	"net/http"
	"productos-api/src/products/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIdController struct {
	getProductByIdUseCase *application.GetProductByIdUseCase
}

func NewGetProductByIdController(getUseCase *application.GetProductByIdUseCase) *GetProductByIdController {
	return &GetProductByIdController{
		getProductByIdUseCase: getUseCase,
	}
}

func (gpc *GetProductByIdController) GetProductById(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, err := gpc.getProductByIdUseCase.Execute(int32(id))

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data": gin.H{
			"type": "products",
			"id":   product.Id,
			"attributes": gin.H{
				"name":     product.Name,
				"price":    product.Price,
				"quantity": product.Quantity,
			},
		},
	}
	g.JSON(http.StatusOK, response)
}
