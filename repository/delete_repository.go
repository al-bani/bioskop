package repository

import (
	"database/sql"
	"bioskop/domain"
)

func DeleteBioskop(db *sql.DB, b domain.Bioskop) (err error) {
	query := "DELETE FROM bioskop WHERE id = $1"

	errs := db.QueryRow(query, b.ID)

	return errs.Err()
}