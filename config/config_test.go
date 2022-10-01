package config

import (
	"go-payment-service/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConection(t *testing.T) {
	config := New("../.env")
	db := NewMysqlDatabase(config)

	con, err := db.DB()
	exception.PanicIfNeeded(err)

	err = con.Ping()

	assert.Nil(t, err)
}

func TestGetValueEnv(t *testing.T) {
	config := New("../.env")
	assert.NotEmpty(t, config.Get("DSN_MYSQL"))
}
