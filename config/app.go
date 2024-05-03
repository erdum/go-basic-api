package config

import (
	"fmt"

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

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if err := envconfig.Process("", appConfig); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return appConfig, nil
}

func GetConfig() *Config {
	return appConfig
}
