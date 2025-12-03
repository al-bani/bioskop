package main

import (
	"bioskop/database"
	"bioskop/routers"
	"fmt"
)

const (
	PORT = ":8080"
)

func init() {
	database.InitDB()
}

func main() {
	router := routers.StartServer()
	router.Run(PORT)

	fmt.Println("Server berjalan di port :", PORT)
}
