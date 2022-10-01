package config

import "os"

type Config interface {
	Get(key string) string
}

type configList struct{}

func (config *configList) Get(key string) string {
	return os.Getenv(key)
}

func New(filesname ...string) Config {
	return &configList{}
}
