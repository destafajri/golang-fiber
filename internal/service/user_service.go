package service

import (
	"github.com/destafajri/golang-fiber/app/model"
	"github.com/destafajri/golang-fiber/app/repository"
)

type UserService interface {
	Register(users *model.RegisterUserPayload) (*model.RegisterUserResponse, error)
	GetData(*model.GetUserPayload) (*model.GetUserResponse , error)
	Login(*model.LoginPayload) (*model.LoginResponse, error)
}

type userServiceimpl struct{
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService{
	return &userServiceimpl{
		userRepository: *userRepository,
	}
}
