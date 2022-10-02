package controllers

import (
	"go-payment-service/models"
	"go-payment-service/services"

	"github.com/gofiber/fiber/v2"
)

type RoomController struct {
	RoomService services.RoomService
}

func NewRoomController(roomService *services.RoomService) RoomController {
	return RoomController{RoomService: *roomService}
}

func (controller *RoomController) Route(app *fiber.App) {
	app.Post("/room", controller.Create)
	app.Get("/room", controller.GetAll)
}

func (controller *RoomController) Create(c *fiber.Ctx) error {
	var request models.Room

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	response := controller.RoomService.Create(request)

	return c.JSON(response)
}

func (controller *RoomController) GetAll(c *fiber.Ctx) error {
	response := controller.RoomService.GetAll()

	return c.JSON(response)
}
