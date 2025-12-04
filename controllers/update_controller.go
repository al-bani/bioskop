package controllers

import (
	"bioskop/database"
	"bioskop/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"bioskop/repository"
	"strconv"
	"bioskop/validator"
)


func UpdateBioskop(ctx *gin.Context) {
	var bioskop domain.Bioskop

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.BindJSON(&bioskop)
    if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal membaca data bioskop",
		})
		return
    }

	bioskop.ID = id

	if !validator.NullCheck(&bioskop) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "nama atau lokasi tidak boleh kosong",
		})
		return
	}
	
	if !validator.RatingCheck(&bioskop) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "rating tidak boleh kurang dari 0 atau lebih dari 10",
		})
		return
	}

	errs := repository.UpdateBioskop(database.DB, bioskop)

	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   errs.Error(),
			"message": "Gagal update data bioskop",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil diupdate",
		"data": bioskop,
	})
}