package routers

import (
	"bioskop/controllers"
	"bioskop/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", middleware.CORSMiddleware(controllers.CreateBioskop))
	router.GET("/bioskop", middleware.CORSMiddleware(controllers.GetAllBioskop))
	router.GET("/bioskop/:id", middleware.CORSMiddleware(controllers.GetBioskop))
	router.PUT("/bioskop/:id", middleware.CORSMiddleware(controllers.UpdateBioskop))
	router.DELETE("/bioskop/:id", middleware.CORSMiddleware(controllers.DeleteBioskop))

	return router
}
