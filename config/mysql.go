package config

import (
	e "go-payment-service/entity"
	"go-payment-service/exception"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase(configuration Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(configuration.Get("DSN_MYSQL")), &gorm.Config{})
	exception.PanicIfNeeded(err)

	db.AutoMigrate(&e.Room{})

	return db
}
