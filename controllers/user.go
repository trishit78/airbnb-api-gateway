package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserByID (w http.ResponseWriter, r *http.Request){
	fmt.Println("creating user in userController")
	//uc.UserService.CreateUser()
	uc.UserService.GetUserByID()
	w.Write([]byte("User fetching endpoint done"))
}



func (uc *UserController) CreateUser (w http.ResponseWriter, r *http.Request){
	fmt.Println("creating user in userController")
	uc.UserService.CreateUser()
	w.Write([]byte("User fetching endpoint done"))
}


func (uc *UserController) LoginUser (w http.ResponseWriter, r *http.Request){
	fmt.Println("LoginUser called in userController")
	uc.UserService.LoginUser()
	
	w.Write([]byte("User login endpoint done"))
}
