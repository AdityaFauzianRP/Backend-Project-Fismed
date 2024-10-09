package model

type Price struct {
	ID             int    `json:"id"`
	NamarumahSakit string `json:"nama_Rumah_Sakit"`
	Kode           string `json:"kode"`
	Variable       string `json:"variable"`
	Nama           string `json:"nama"`
	Name           string `json:"name"`
	Diskon         int    `json:"diskon"`
	Price          string `json:"price"`
	Added          string `json:"added"`
}

type Set struct {
	Input []Price `json:"input"`
}
