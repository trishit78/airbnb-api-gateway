package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte (plainPassword),bcrypt.DefaultCost)
	if err != nil{
		fmt.Println("Error hashing password:",err)
		return "",err
	}
	return string(hash),nil
}



func CheckPasswordHash(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}







// package services

// import (
// 	env "AuthInGo/config/env"
// 	db "AuthInGo/db/repositories"
// 	"AuthInGo/utils"
// 	"fmt"

// 	"github.com/golang-jwt/jwt/v5"
// )

// type UserService interface {
// 	GetUserById() error
// 	CreateUser() error
// 	LoginUser() (string, error)
// }

// type UserServiceImpl struct {
// 	userRepository db.UserRepository
// }

// func NewUserService(_userRepository db.UserRepository) UserService {
// 	return &UserServiceImpl{
// 		userRepository: _userRepository,
// 	}
// }

// func (u *UserServiceImpl) GetUserById() error {
// 	fmt.Println("Fetching user in UserService")
// 	u.userRepository.GetByID()
// 	return nil
// }

// func (u *UserServiceImpl) CreateUser() error {
// 	fmt.Println("Creating user in UserService")
// 	password := "example_password"
// 	hashedPassword, err := utils.HashPassword(password)
// 	if err != nil {
// 		return err
// 	}
// 	u.userRepository.Create(
// 		"username_example_1",
// 		"user1@example.com",
// 		hashedPassword,
// 	)
// 	return nil
// }
