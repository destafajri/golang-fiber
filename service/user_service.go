package service

import (
	"github.com/destafajri/golang-fiber/model"
	"github.com/destafajri/golang-fiber/repository"
)

type UserService interface {
	Register(users *model.RegisterUserPayload) (*model.RegisterUserResponse, error)
	// GetData(phone string) (model.GetUserResponse)
	// Login(phone string) (model.LoginResponse)
}

type userServiceimpl struct{
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService{
	return &userServiceimpl{
		userRepository: *userRepository,
	}
}
