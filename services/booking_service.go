package services

import (
	"go-payment-service/entity"
	"go-payment-service/models"
	"go-payment-service/repository"
)

type BookingService interface {
	Booking(request models.Booking) (response models.Booking)
}

type BookingServiceImpl struct {
	BookingRepository repository.BookingRepository
}

func NewBookingServiceImpl(bookingRepository *repository.BookingRepository) BookingService {
	return &BookingServiceImpl{BookingRepository: *bookingRepository}
}

func (service *BookingServiceImpl) Booking(request models.Booking) (response models.Booking) {
	booking := entity.Booking{
		RoomID: request.RoomID,
		Price:  request.Price,
	}

	service.BookingRepository.Booking(booking)

	response = models.Booking{
		RoomID: booking.RoomID,
		Price:  booking.Price,
	}

	return response
}
