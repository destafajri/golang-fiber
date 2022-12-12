package service

import (
	_ "strconv"

	"github.com/destafajri/golang-fiber/entity"
	"github.com/destafajri/golang-fiber/helper"
	"github.com/destafajri/golang-fiber/model"
	"github.com/destafajri/golang-fiber/validation"
	"github.com/google/uuid"
)

func(user userServiceimpl) Register(request *model.RegisterUserPayload) (*model.RegisterUserResponse, error) {
	//validation input phone number
	phone, err := validation.Phonenumber(request.Phone)
	if !phone && err != nil{
		return nil, err
	}

	//generate password
	pass := helper.GeneratePassword(4,1,1,1)

	input := entity.UserEntity{
		ID: uuid.New().String(),
		Name: request.Name,
		Phone: request.Phone,
		Role:request.Role,
		Password: pass,
	}
	
	err = user.userRepository.Register(&input)
	if err != nil {
		return nil, err
	}

	response := model.RegisterUserResponse{
		ID: input.ID,
		Name: input.Name,
		Phone: input.Phone,
		Role: input.Role,
		Password: input.Password,
	}
	return &response, nil
}