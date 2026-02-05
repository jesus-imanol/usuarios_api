package routers

import (
	"productos-api/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, registerController *controllers.RegisterUserController, updateController *controllers.UpdateUserController, listUserController *controllers.ListUserController, deleteController *controllers.DeleteUserController, loginUserController *controllers.LoginUserController, getUserByIdController *controllers.GetUserByIdController, uploadPictureController *controllers.UploadPictureUserController) {
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", registerController.RegisterUser)
		v1.POST("/login", loginUserController.LoginUser)
		v1.GET("/:id", getUserByIdController.GetUserByID)
		v1.GET("/", listUserController.GetAllUsers)
		v1.PUT("/:id", updateController.UpdateUser)
		v1.DELETE("/:id", deleteController.DeleteUser)
		v1.PUT("/upload-picture/:id", uploadPictureController.UpdatePictureUser)
	}
}
