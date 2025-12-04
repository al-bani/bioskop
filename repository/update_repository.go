package repository

import (
	"bioskop/domain"
	"database/sql"
)

func UpdateBioskop(db *sql.DB, b domain.Bioskop) (err error) {
	query := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4"

	errs := db.QueryRow(query, b.Nama, b.Lokasi, b.Rating, b.ID)

	return errs.Err()
}