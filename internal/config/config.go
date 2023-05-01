package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	DBSSLMode  string

	AppPort string
}

func InitConfig() Config {
	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE"),

		AppPort: os.Getenv("APP_PORT"),
	}
}

func InitTests() Config {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_TEST_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE"),

		AppPort: os.Getenv("APP_PORT"),
	}
}
