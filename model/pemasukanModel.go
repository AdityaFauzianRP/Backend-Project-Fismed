package model

type Pemasukan struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Nominal string `json:"nominal"`
	Amount  string `json:"amount"`
	Pajak   string `json:"pajak"`
	Tanggal string `json:"tanggal"`
	Status  string `json:"status"`
	Piid    string `json:"pi_id"`
	Divisi  string `json:"divisi"`
}

type Pengeluaran struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Nominal string `json:"nominal"`
	Amount  string `json:"amount"`
	Pajak   string `json:"pajak"`
	Tanggal string `json:"tanggal"`
	Status  string `json:"status"`
	Piid    string `json:"pi_id"`
	Divisi  string `json:"divisi"`
}
