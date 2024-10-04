package model

type ItemPI struct {
	Kode        string `json:"kode"`
	Kat         string `json:"kat"`
	NamaBarang  string `json:"nama_barang"`
	Quantity    string `json:"quantity"`
	HargaSatuan string `json:"harga_satuan"`
	Discount    int    `json:"discount"`
	Variable    string `json:"variable"`
	Amount      string `json:"amount"`
	Gudang      string `json:"gudang"`
	Price       string `json:"price"`
}

type DeletedItem struct {
	Kat string `json:"kat"`
}

type ProformaInvoice struct {
	IDDivisi        string        `json:"id_divisi"`
	RumahSakit      string        `json:"rumah_sakit"`
	Alamat          string        `json:"alamat"`
	JatuhTempo      string        `json:"jatuh_tempo"`
	NamaDokter      string        `json:"nama_dokter"`
	NamaPasien      string        `json:"nama_pasien"`
	RM              string        `json:"rm"`
	IDRumahSakit    string        `json:"id_rumah_sakit"`
	TanggalTindakan string        `json:"tanggal_tindakan"`
	TanggalPI       string        `json:"tanggal"`
	NomorInvoice    string        `json:"nomor_invoice"`
	NomorSI         string        `json:"nomor_si"`
	Pajak           string        `json:"pajak"`
	Subtotal        string        `json:"subtotal"`
	Total           string        `json:"total"`
	RP_sub_total    string        `json:"RP_sub_total"`
	Item            []ItemPI      `json:"item"`
	ItemDeleted     []DeletedItem `json:"item_deleted"`
}
