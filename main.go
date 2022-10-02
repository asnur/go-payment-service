package main

import (
	"go-payment-service/config"
	"go-payment-service/controllers"
	"go-payment-service/repository"
	"go-payment-service/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	roomRepository := repository.NewRoomRepositoryImpl(database)

	roomService := services.NewRoomServiceImpl(&roomRepository)

	roomController := controllers.NewRoomController(&roomService)

	roomController.Route(app)

	app.Listen(":3000")
}
