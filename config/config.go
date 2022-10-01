package config

import (
	"go-payment-service/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configList struct{}

func (config *configList) Get(key string) string {
	return os.Getenv(key)
}

func New(filesname ...string) Config {
	err := godotenv.Load(filesname...)
	exception.PanicIfNeeded(err)

	return &configList{}
}
