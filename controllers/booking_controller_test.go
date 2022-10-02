package controllers

import (
	"bytes"
	"encoding/json"
	"go-payment-service/models"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookingController(t *testing.T) {
	var bookingRequest models.Booking

	bookingRequest.RoomID = 1
	bookingRequest.Price = 100000

	requestBody, _ := json.Marshal(bookingRequest)

	request := httptest.NewRequest("POST", "/booking", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}
