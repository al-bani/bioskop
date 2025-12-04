package repository

import (
	"bioskop/domain"
	"database/sql"
)

func GetAllBioskop(db *sql.DB) (res []domain.Bioskop, err error) {
	query := "SELECT * FROM bioskop"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b domain.Bioskop

		err = rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)

		if err != nil {
			return nil, err
		}

		res = append(res, b)
	}

	return res, nil
}