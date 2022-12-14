package product

import (
	"github.com/destafajri/golang-fiber/middlewares"
	"github.com/gofiber/fiber/v2"
)

func (controller *ProductController, ) Route(app *fiber.App) {
	// Setup Versioning Route
	v1 := app.Group("/v1", middlewares.New())
	
	v1.Post("/api/products", controller.Create)
	v1.Get("/api/products", controller.List)
}