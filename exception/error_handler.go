package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/destafajri/golang-fiber/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:   fiber.StatusUnprocessableEntity,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
