package controllers

import (
	"net/http"
	"productos-api/src/users/application"
	"productos-api/src/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	loginUserUseCase *application.LoginUserUseCase
}

func NewLoginUserController(loginUseCase *application.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{loginUserUseCase: loginUseCase}
}

func (luc *LoginUserController) LoginUser(g *gin.Context) {
	var user entities.User
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := luc.loginUserUseCase.Execute(user.Email, user.PasswordHash)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := GenerateJWT(*foundUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.Header("Authorization", "Bearer "+token)

	response := gin.H{
		"data": gin.H{
			"type": "users",
			"id":   foundUser.Id,
			"attributes": gin.H{
				"full_name": foundUser.FullName,
				"email":     foundUser.Email,
			},
		},
	}
	g.JSON(http.StatusOK, response)
}
