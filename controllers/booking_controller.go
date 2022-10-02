package controllers

import (
	"go-payment-service/models"
	"go-payment-service/services"

	"github.com/gofiber/fiber/v2"
)

type BookingController struct {
	BookingService services.BookingService
}

func NewBookingController(bookingService *services.BookingService) BookingController {
	return BookingController{BookingService: *bookingService}
}

func (controller *BookingController) Route(app *fiber.App) {
	app.Post("/booking", controller.Booking)
}

func (controller *BookingController) Booking(c *fiber.Ctx) error {
	var request models.Booking

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	response := controller.BookingService.Booking(request)

	return c.JSON(response)
}
