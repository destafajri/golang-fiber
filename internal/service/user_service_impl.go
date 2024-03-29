package service

import (
	"errors"
	_ "strconv"

	"github.com/destafajri/golang-fiber/internal/entity"
	"github.com/destafajri/golang-fiber/internal/model"
	"github.com/destafajri/golang-fiber/internal/validation"
	"github.com/destafajri/golang-fiber/helper"
	"github.com/google/uuid"
)

func (user userServiceimpl) Register(request *model.RegisterUserPayload) (*model.RegisterUserResponse, error) {
	//validation input phone number
	phone, err := validation.Phonenumber(request.Phone)
	if !phone && err != nil {
		return nil, err
	}

	//generate password
	pass := helper.GeneratePassword(4, 1, 1, 1)

	input := entity.UserEntity{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Phone:    request.Phone,
		Role:     request.Role,
		Password: pass,
	}

	err = user.userRepository.Register(&input)
	if err != nil {
		return nil, err
	}

	response := model.RegisterUserResponse{
		ID:       input.ID,
		Name:     input.Name,
		Phone:    input.Phone,
		Role:     input.Role,
		Password: input.Password,
	}
	return &response, nil
}

func (user userServiceimpl) GetData(input *model.GetUserPayload) (*model.GetUserResponse, error) {
	//validation input phone number
	phone, err := validation.Phonenumber(input.Phone)
	if !phone && err != nil {
		return nil, err
	}

	users, err := user.userRepository.GetData(input.Phone)
	if err != nil {
		return nil, err
	}

	return (*model.GetUserResponse)(users), nil
}

func (user userServiceimpl) Login(input *model.LoginPayload) (*model.LoginResponse, error) {
	//validation input phone number
	phone, err := validation.Phonenumber(input.Phone)
	if !phone && err != nil {
		return nil, err
	}

	users, err := user.userRepository.Login(input.Phone)
	if err != nil {
		return nil, err
	}

	//password check
	if input.Password != users.Password {
		return nil, errors.New("You're Unauthorized")
	}

	//Generate JWT
	token, err := helper.GenerateJwtToken(users.Name, users.Phone, users.Role)
	if err != nil {
		return nil, err		
	}

	resp := model.LoginResponse{
		Token: token,
	}

	return &resp, nil
}
