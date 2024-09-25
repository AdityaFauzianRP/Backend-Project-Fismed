package model

type Price struct {
	NamarumahSakit string `json:"nama_Rumah_Sakit"`
	Kode           string `json:"kode"`
	Variable       string `json:"variable"`
	Nama           string `json:"nama"`
	Diskon         int    `json:"diskon"`
	Price          string `json:"price"`
}

type Set struct {
	Input []Price `json:"input"`
}
