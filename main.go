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

	//Config
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	// Room Logic
	roomRepository := repository.NewRoomRepositoryImpl(database)
	roomService := services.NewRoomServiceImpl(&roomRepository)
	roomController := controllers.NewRoomController(&roomService)
	roomController.Route(app)

	// Booking Logic
	bookingRepository := repository.NewBookingRepositoryImpl(database)
	bookingService := services.NewBookingServiceImpl(&bookingRepository)
	bookingController := controllers.NewBookingController(&bookingService)
	bookingController.Route(app)

	app.Listen(":3000")
}
