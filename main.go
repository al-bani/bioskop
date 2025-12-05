package main

import (
	"bioskop/config"
	"bioskop/database"
	"bioskop/routers"
	"fmt"
)

func init() {
	config.LoadConfig()
	database.InitDB()
}

func main() {
	router := routers.StartServer()
	router.Run(":8080")

	fmt.Println("Server berjalan di port berapa bang :", config.PORT)
}
