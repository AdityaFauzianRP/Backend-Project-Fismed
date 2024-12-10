package model

type Stock struct {
	Id               int    `json:"id"`
	Variable         string `json:"variable"`
	Nama             string `json:"nama"`
	Qty              int    `json:"qty"`
	Harga            string `json:"harga"`
	Kode             string `json:"kode"`
	NamaGudang       string `json:"namaGudang,omitempty"`
	Lots             string `json:"lots"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	KeteranganBarang string `json:"keterangan_barang"`
}

type BarangTerjual struct {
	NamaBarang string `json:"nama_barang"`
	Qty        string `json:"quantity"`
	Tanggal    string `json:"tanggal_terjual"`
	Customer   string `json:"customer"`
}

type Gudang struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama_gudang"`
	Lokasi string `json:"alamat_gudang"`
}
