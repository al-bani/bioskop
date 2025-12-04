package repository

import (
	"bioskop/domain"
	"database/sql"
)

func CreateBioskop(db *sql.DB, b domain.Bioskop) (err error) {
	query := "INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3)"

	errs := db.QueryRow(query, b.Nama, b.Lokasi, b.Rating)

	return errs.Err()
}