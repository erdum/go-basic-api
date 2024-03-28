package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug               bool
	Port                string
	Url                 string
	FirebaseCredentials string
}

var appConfig = &Config{
	Debug:               true,
	Port:                "8000",
	Url:                 "http://localhost",
	FirebaseCredentials: "firebase_credentials.json",
}

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
