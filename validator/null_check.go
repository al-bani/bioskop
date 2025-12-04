package validator

import (
	"bioskop/domain"
)

func NullCheck(b *domain.Bioskop) bool {
	if b.Nama == "" || b.Lokasi == ""{
		return false
	}

	return true
}

func RatingCheck(b *domain.Bioskop) bool {
	if b.Rating < 0 || b.Rating > 10 {
		return false
	}
	return true
}