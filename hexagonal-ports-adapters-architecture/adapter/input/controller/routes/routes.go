package routes

import (
	"golang_hexagonal_architecture/adapter/input/controller"

	"github.com/julienschmidt/httprouter"
)

func InitRoutes(r *httprouter.Router, userController controller.UserControllerInterface) {
	r.POST("/user", userController.CreateUser)
}
