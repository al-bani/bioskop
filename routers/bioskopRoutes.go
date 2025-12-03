package routers

import (
	"bioskop/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)

	return router
}