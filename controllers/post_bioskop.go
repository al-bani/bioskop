package controllers

import (
	"bioskop/database"
	"bioskop/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"bioskop/validator"
	"errors"
)

func CreateBioskop(ctx *gin.Context) {
	var newBioskop model.Bioskop

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if !validator.NullCheck(&newBioskop) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("nama dan lokasi tidak boleh kosong"))
		return
	}

	if newBioskop.Rating < 0 || newBioskop.Rating > 10 {
		newBioskop.Rating = 0
	}

	model.ListBioskop = append(model.ListBioskop, newBioskop)

	query := "INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, newBioskop.Nama, newBioskop.Lokasi, newBioskop.Rating)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil dibuat",
		"data": newBioskop,
	})
	
	panic("Kesalahan Server")

}
