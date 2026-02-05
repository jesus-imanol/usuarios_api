package dependenciesuser

import (
	"productos-api/src/users/application"
	"productos-api/src/users/infraestructure/adapters"
	"productos-api/src/users/infraestructure/controllers"
	"productos-api/src/users/infraestructure/routers"

	"github.com/gin-gonic/gin"
)

func InitUsers(r *gin.Engine) {
	// Inicializar adaptador MySQL
	mysql, err := adapters.NewMySQL()
	if err != nil {
		panic(err)
	}

	// Inicializar casos de uso
	registerUseCase := application.NewRegisterUserUseCase(mysql)
	updateUseCase := application.NewUpdateUserUseCase(mysql)
	listUserUseCase := application.NewListUserUseCase(mysql)
	deleteUserUseCase := application.NewDeleteUserUseCase(mysql)
	loginUseCase := application.NewLoginUserUseCase(mysql)
	getByIdUseCase := application.NewGetUserById(mysql)
	uploadPictureUseCase := application.NewUploadPictureUserUseCase(mysql)

	// Inicializar controladores
	registerUser_controller := controllers.NewRegisterUserController(registerUseCase)
	updateUser_controller := controllers.NewUpdateUserController(updateUseCase)
	listUser_controller := controllers.NewListUserController(listUserUseCase)
	deleteUser_controller := controllers.NewDeleteUserController(deleteUserUseCase)
	loginUser_controller := controllers.NewLoginUserController(loginUseCase)
	getUserById_controller := controllers.NewGetUserByIDController(getByIdUseCase)
	uploadPictureUser_controller := controllers.NewUploadPictureUserController(uploadPictureUseCase)

	// Inicializar router
	routers.UserRoutes(r, registerUser_controller, updateUser_controller, listUser_controller, deleteUser_controller, loginUser_controller, getUserById_controller, uploadPictureUser_controller)
}
