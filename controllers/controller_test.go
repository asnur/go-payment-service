package controllers

import (
	"go-payment-service/config"
	"go-payment-service/repository"
	"go-payment-service/services"

	"github.com/gofiber/fiber/v2"
)

func createTestApp() *fiber.App {
	app := fiber.New()
	roomController.Route(app)
	return app
}

var configuration = config.New("../.env")
var database = config.NewMysqlDatabase(configuration)
var roomRepository = repository.NewRoomRepositoryImpl(database)
var roomService = services.NewRoomServiceImpl(&roomRepository)

var roomController = NewRoomController(&roomService)

var app = createTestApp()
