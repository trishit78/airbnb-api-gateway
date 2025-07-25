package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController)Router{
	return &UserRouter{
		userController: _userController ,
	}
}


func (ur *UserRouter) Register(r chi.Router){
	r.Get("/profile", ur.userController.GetUserByID)
	r.Post("/signup", ur.userController.CreateUser)
	r.Post("/login", ur.userController.LoginUser)
}