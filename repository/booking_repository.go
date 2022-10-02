package repository

import (
	"go-payment-service/entity"

	"gorm.io/gorm"
)

type BookingRepository interface {
	Booking(request entity.Booking)
	List() (response []entity.Booking)
}

type BookingRepositoryImpl struct {
	Query *gorm.DB
}

func NewBookingRepositoryImpl(query *gorm.DB) BookingRepository {
	return &BookingRepositoryImpl{Query: query}
}

func (repository *BookingRepositoryImpl) Booking(request entity.Booking) {
	repository.Query.Create(&request)
}

func (repository *BookingRepositoryImpl) List() (response []entity.Booking) {
	repository.Query.Find(&response)
	return response
}
