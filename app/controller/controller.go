package controller

import (
	"github.com/destafajri/golang-fiber/app/controller/product"
	"github.com/destafajri/golang-fiber/app/controller/user"
	"github.com/destafajri/golang-fiber/app/repository"
	"github.com/destafajri/golang-fiber/app/service"
	"github.com/destafajri/golang-fiber/config"
	"github.com/destafajri/golang-fiber/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Controller(){
	// Setup Configuration
	configuration 	:= config.New()
	databaseMongo 	:= config.NewMongoDatabase(configuration)
	databasePostgre := config.NewPostgreDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(databaseMongo)
	userRepository := repository.NewUserRepository(databasePostgre)

	// Setup Service
	productService := service.NewProductService(&productRepository)
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	productController := product.NewProductController(&productService)
	userController := user.NewUserController(&userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	productController.Route(app)
	userController.Route(app)

	// Start App
	err := app.Listen("0.0.0.0:9000")
	exception.PanicIfNeeded(err)
}