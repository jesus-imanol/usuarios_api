package controllers

import (
	"net/http"
	"productos-api/src/products/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	deleteProductUseCase *application.DeleteProductUseCase
}

func NewDeleteProductController(deleteUseCase *application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{
		deleteProductUseCase: deleteUseCase,
	}
}

func (dpc *DeleteProductController) DeleteProduct(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dpc.deleteProductUseCase.Execute(int32(id))

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Producto eliminado exitosamente",
	})
}
