package controllers

import (
	"fmt"
	"net/http"
	"productos-api/src/users/application"

	"github.com/gin-gonic/gin"
)

type ListUserController struct {
	listUserUseCase *application.ListUserUseCase
}

func NewListUserController(useCase *application.ListUserUseCase) *ListUserController {
	return &ListUserController{listUserUseCase: useCase}
}

func (lu *ListUserController) GetAllUsers(g *gin.Context) {
	list_users, err := lu.listUserUseCase.Execute()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, user := range list_users {
		userResponse := gin.H{
			"type": "users",
			"id":   user.Id,
			"attributes": gin.H{
				"full_name": user.FullName,
				"email":     user.Email,
			},
		}
		response = append(response, userResponse)
	}

	if len(list_users) > 0 {
		g.JSON(http.StatusOK, gin.H{"data": response})
	} else {
		fmt.Println("Users:", len(list_users))
		g.JSON(http.StatusOK, gin.H{
			"data":    len(list_users),
			"message": "No se encontraron users",
			"type":    "users",
		})
	}
}
