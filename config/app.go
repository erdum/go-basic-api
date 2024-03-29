package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug    bool   `default:"true"`
	Port     string `default:"8000"`
	Url      string `default:"http://localhost"`
	Firebase struct {
		ProjectId   string `split_words:"true"`
		Credentials string
	}
}

var appConfig = &Config{}

func LoadConfig() *Config {
	godotenv.Load()

	if err := envconfig.Process("", appConfig); err != nil {
		panic(err)
	}

	return appConfig
}

func GetConfig() *Config {
	return appConfig
}
