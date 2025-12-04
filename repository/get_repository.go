package repository

import (
	"bioskop/domain"
	"database/sql"
)

func GetBioskop(db *sql.DB, b domain.Bioskop) (res domain.Bioskop, err error) {
	query := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1"

	err = db.QueryRow(query, b.ID).Scan(
		&res.ID,
		&res.Nama,
		&res.Lokasi,
		&res.Rating,
	)
	
	if err != nil {
		return domain.Bioskop{}, err
	}
	
	return res, err
}