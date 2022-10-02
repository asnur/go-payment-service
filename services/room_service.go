package services

import (
	"go-payment-service/entity"
	"go-payment-service/models"
	"go-payment-service/repository"
)

type RoomService interface {
	Create(request models.Room) (response models.Room)
	GetAll() (response []models.Room)
}

func NewRoomServiceImpl(roomRepository *repository.RoomRepository) RoomService {
	return &RoomServiceImpl{RoomRepository: *roomRepository}

}

type RoomServiceImpl struct {
	RoomRepository repository.RoomRepository
}

func (service *RoomServiceImpl) Create(request models.Room) (response models.Room) {
	room := entity.Room{
		Name:  request.Name,
		Price: request.Price,
	}

	service.RoomRepository.Insert(room)

	response = models.Room{
		Name:  room.Name,
		Price: room.Price,
	}

	return response
}

func (service *RoomServiceImpl) GetAll() (response []models.Room) {
	rooms := service.RoomRepository.FindAll()

	for _, room := range rooms {
		response = append(response, models.Room{
			Name:  room.Name,
			Price: room.Price,
		})
	}

	return response
}
