package controllers

import (
	"bioskop/database"
	"bioskop/domain"
	"bioskop/repository"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetBioskop(ctx *gin.Context) {
	var bioskop domain.Bioskop

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.BindJSON(&bioskop)
	if err != nil {
		panic(err)
	}

	bioskop.ID = id

	bioskop, errs := repository.GetBioskop(database.DB, bioskop)

	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   errs.Error(),
			"message": "Gagal mengambil data bioskop",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil diambil",
		"data":    bioskop,
	})
}
