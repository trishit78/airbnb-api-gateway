package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	GetUserByID() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository  db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

// func (u *UserServiceImpl) CreateUser() error {
// 	fmt.Println("Creating user in UserService")
// 	u.userRepository.Create()
// 	return nil
// }


func (u *UserServiceImpl) GetUserByID() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.Create()
	return nil
}