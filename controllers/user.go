package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
	"encoding/json"
    
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
	user ,err :=uc.UserService.GetUserByID(4)
	 
	fmt.Println("user is",user)
	fmt.Println("err is",err)
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

	response:= fmt.Sprintf("user got successfully: %+v", user)
	 json.NewEncoder(w).Encode(response)
}



func (uc *UserController) CreateUser (w http.ResponseWriter, r *http.Request){
	fmt.Println("creating user in userController")
	uc.UserService.CreateUser()
	w.Write([]byte("User fetching endpoint done"))

	 w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	response:= "user created successfully"
	 json.NewEncoder(w).Encode(response)
}


func (uc *UserController) LoginUser (w http.ResponseWriter, r *http.Request){
	fmt.Println("LoginUser called in userController")
	uc.UserService.LoginUser()
	
	 w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

	response:= "user logged in successfully"
	 json.NewEncoder(w).Encode(response)
}
