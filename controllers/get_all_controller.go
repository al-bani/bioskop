package controllers

import (
	"bioskop/database"
	"bioskop/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"bioskop/repository"
)


func GetAllBioskop(ctx *gin.Context) {
	var err error

	domain.ListBioskop, err = repository.GetAllBioskop(database.DB)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal menampilkan data bioskop",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List Bioskop berhasil diambil",
		"data": domain.ListBioskop,
	})
}
