package models

type Booking struct {
	RoomID int     `json:"room_id"`
	Price  float64 `json:"price"`
}
