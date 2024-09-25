package model

type Stock struct {
	Variable   string `json:"variable"`
	Nama       string `json:"nama"`
	Qty        int    `json:"qty"`
	Harga      string `json:"harga"`
	Kode       string `json:"kode"`
	NamaGudang string `json:"namaGudang,omitempty"`
}
