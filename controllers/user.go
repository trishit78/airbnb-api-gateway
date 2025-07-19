package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"encoding/json"
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
	// create a dto for payload

	//readjson and give error

	//validate payload


	user ,err :=uc.UserService.GetUserByID(4)
	 
	//deal with this err

	//success response
	
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



func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging in", jsonErr)
		return
	}

	fmt.Println("Payload received:", payload)

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)

}
