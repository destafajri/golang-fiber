package user

import (
	"github.com/destafajri/golang-fiber/model"
	"github.com/destafajri/golang-fiber/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.RegisterUserPayload
	
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.Register(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "errors",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   response,
	})
}
