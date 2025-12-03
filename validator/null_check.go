package validator

import (
	"bioskop/model"
)

func NullCheck(b *model.Bioskop) bool {
	if b.Nama == "" || b.Lokasi == ""{
		return false
	}

	return true
}