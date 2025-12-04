package controllers

import (
	"bioskop/database"
	"bioskop/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"bioskop/validator"
	"bioskop/repository"
)

func CreateBioskop(ctx *gin.Context) {
	var newBioskop domain.Bioskop

	if errs := ctx.ShouldBindJSON(&newBioskop); errs != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   errs.Error(),
			"message": "tidak dapat membaca data bioskop",
		})
		return
	}

	if !validator.NullCheck(&newBioskop) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "nama tidak boleh kosong",
		})
		return
	}

	if !validator.RatingCheck(&newBioskop) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "rating tidak boleh kurang dari 0 atau lebih dari 10",
		})
		return
	}

	domain.ListBioskop = append(domain.ListBioskop, newBioskop)

	err := repository.CreateBioskop(database.DB, newBioskop)
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Gagal membuat data bioskop",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil dibuat",
		"data": newBioskop,
	})

}
