package controllers

import (
	"bioskop/database"
	"bioskop/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"bioskop/repository"
	"strconv"
)

func DeleteBioskop(ctx *gin.Context) {
	var bioskop domain.Bioskop

	id, _ := strconv.Atoi(ctx.Param("id"))

    err := ctx.BindJSON(&bioskop)
    if err != nil {
       panic(err)
    }

    bioskop.ID = id

	errs := repository.DeleteBioskop(database.DB, bioskop)

	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   errs.Error(),
			"message": "Gagal menghapus data bioskop",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil dihapus",
		"data": bioskop,
	})
}