package repository

import (
	"go-payment-service/entity"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Insert(request entity.Room)
	FindAll() (response []entity.Room)
}

type RoomRepositoryImpl struct {
	Query *gorm.DB
}

func NewRoomRepositoryImpl(query *gorm.DB) RoomRepository {
	return &RoomRepositoryImpl{Query: query}
}

func (repository *RoomRepositoryImpl) Insert(request entity.Room) {
	repository.Query.Create(&request)
}

func (repository *RoomRepositoryImpl) FindAll() (response []entity.Room) {
	repository.Query.Find(&response)
	return response
}
