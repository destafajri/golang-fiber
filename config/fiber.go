package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/destafajri/golang-fiber/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
