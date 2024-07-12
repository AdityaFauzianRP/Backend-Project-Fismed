package model

import "time"

type Pemasukan struct {
	Id      int       `json:"id,omitempty"`
	Nama    string    `json:"nama,omitempty"`
	Nominal string    `json:"nominal,omitempty"`
	Amount  string    `json:"amount,omitempty"`
	Tanggal time.Time `json:"tanggal,omitempty"`
}
