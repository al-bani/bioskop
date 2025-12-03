package model

type Bioskop struct {
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}

var ListBioskop []Bioskop