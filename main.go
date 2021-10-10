package main

import (
	"belajar-mongo/controller"
	"belajar-mongo/pkg/configuration"
	"belajar-mongo/pkg/database"
	"belajar-mongo/pkg/exception"
	"belajar-mongo/repository"
	"belajar-mongo/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := configuration.NewConfigurationImpl(".env")
	db := database.NewMongoDatabase(config)
	validate := validator.New()

	newsRepository := repository.NewNewsRepositoryImpl(db)
	newsService := service.NewNewsServiceImpl(newsRepository, validate)
	newsController := controller.NewNewsController(newsService)

	app := fiber.New(configuration.NewFiberConfig())
	app.Use(recover.New())

	app.Static("/uploads", "./uploads")
	newsController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
