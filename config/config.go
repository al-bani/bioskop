package config

import (
	"os"
	"github.com/joho/godotenv"
)

var (
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PSWD string
	DB_NAME string
	PORT    string
)

func LoadConfig() {
	godotenv.Load("config/.env")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PSWD = os.Getenv("DB_PSWD")
	DB_NAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("PORT")
}