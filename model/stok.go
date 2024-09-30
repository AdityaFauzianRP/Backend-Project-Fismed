package model

type Stock struct {
	Id         int    `json:"id"`
	Variable   string `json:"variable"`
	Nama       string `json:"nama"`
	Qty        int    `json:"qty"`
	Harga      string `json:"harga"`
	Kode       string `json:"kode"`
	NamaGudang string `json:"namaGudang,omitempty"`
	Name       string `json:"name"`
	Price      string `json:"price"`
}
