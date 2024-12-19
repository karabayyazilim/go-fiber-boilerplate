package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

var (
	AppPort    string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	DbSSLMode  string
	DbTimeZone string
)

func init() {
	currentWorkDirectory, _ := os.Getwd()

	err := godotenv.Load(string(currentWorkDirectory) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}
}

type Environment struct {
	AppPort    string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	DbSSLMode  string
	DbTimeZone string
}

var app = &Environment{}

func Env() *Environment {
	app.AppPort = os.Getenv("APP_PORT")
	app.DbHost = os.Getenv("DB_HOST")
	app.DbPort = os.Getenv("DB_PORT")
	app.DbName = os.Getenv("DB_NAME")
	app.DbUser = os.Getenv("DB_USER")
	app.DbPassword = os.Getenv("DB_PASSWORD")
	app.DbSSLMode = os.Getenv("DB_SSL_MODE")
	app.DbTimeZone = os.Getenv("DB_TIMEZONE")

	return app
}
