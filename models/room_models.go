package models

type Room struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateRoom struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
