package database

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq" 
)

var DB *sql.DB

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PSWD = "root"
	DB_NAME = "sanber"
)

func InitDB() {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PSWD, DB_NAME)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Berhasil terhubung ke database")
}
