package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-payment-service/models"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	var roomRequest models.Room
	roomRequest.Name = "Room 1"
	roomRequest.Price = 100

	requestBody, _ := json.Marshal(roomRequest)
	request := httptest.NewRequest("POST", "/room", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}

func TestGetAllRoom(t *testing.T) {
	request := httptest.NewRequest("GET", "/room", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	fmt.Println(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}
