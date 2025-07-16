package services

import (
	_ "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"

	_ "github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByID() error
	CreateUser() error
	LoginUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "example_password"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.userRepository.Create(
		"username_example",
		"username@example.com",
		hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) GetUserByID() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.DeleteByID(1)
	return nil
}
func (u *UserServiceImpl) LoginUser() error {
	// Pre-requisite: This function will be given email and password as parameter, which we can hardcode for now
	email := "username@example.com"
	password := "example_password"
	// Step 1. Make a repository call to get the user by email
	matchUser, err := u.userRepository.GetByEmail(email)
	// Step 2. If user exists, or not. If not exists, return error
	if err != nil {
		fmt.Println("user not matched")
	}
	// Step 3. If user exists, check the password using utils.CheckPasswordHash
	res := utils.CheckPasswordHash(password, matchUser.Password)
	// Step 4. If password matches, print a JWT token, else return error saying password does not match

	if !res {
		fmt.Println("password does not match")
		return nil
	}
	fmt.Println("login success")
	// payload:= jwt.MapClaims{
	// 	"email":matchUser.Email,
	// 	"id":matchUser.Id,
	// }

	// token:= jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	// tokenString,err :=token.SignedString([]byte(env.GetString("JWT_SECRET","TOKEN")))

	// if err!=nil{
	// 	fmt.Println("Generated JWT Token",tokenString)
	// }
	return nil
}
