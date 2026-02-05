package controllers

import (
	"net/http"
	"productos-api/src/users/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserByIdController struct {
	getUserByIdUseCase *application.GetUserByIdUseCase
}

func NewGetUserByIDController(useCase *application.GetUserByIdUseCase) *GetUserByIdController {
	return &GetUserByIdController{getUserByIdUseCase: useCase}
}

func (gubi *GetUserByIdController) GetUserByID(g *gin.Context) {
	idParam := g.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	idGet := int32(id)
	user, err := gubi.getUserByIdUseCase.Execute(idGet)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{
		"data": gin.H{
			"type": "users",
			"id":   user.Id,
			"attributes": gin.H{
				"full_name":        user.FullName,
				"email":            user.Email,
				"profile_picture":  user.ProfilePicture,
				"gender":           user.Gender,
				"city":             user.City,
				"state":            user.State,
				"status_message":   user.StatusMessage,
				"match_preference": user.MatchPreference,
				"interests":        user.Interests,
			},
		},
	}
	g.JSON(http.StatusOK, response)
}
