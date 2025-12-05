package config

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
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
	godotenv.Load()

	DB_HOST = os.Getenv("PGHOST")
	DB_PORT = os.Getenv("PGPORT")
	DB_USER = os.Getenv("PGUSER")
	DB_PSWD = os.Getenv("PGPASSWORD")
	DB_NAME = os.Getenv("PGDATABASE")
	PORT = os.Getenv("PORT")

	fmt.Println("Loaded PORT from .env:", PORT)
	

	if PORT == "" {
		fmt.Println("ternyata tidak ada, jadi di set manual")
		PORT = ":8080"
	}
}