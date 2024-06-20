package service

import "time"

type StockBarang struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Total     string    `json:"total"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty"`
}

type RequestID struct {
	ID string `json:"id"`
}

type PerformanceInvoice struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	ItemName      string `json:"item_name"`
	Discount      string `json:"discount"`
	Status        string `json:"status"`
	Divisi        string `json:"divisi"`
	InvoiceNumber int    `json:"invoice_number"`
	PONumber      int    `json:"po_number"`
	SubTotal      string `json:"sub_total"`
	Pajak         string `json:"pajak"`
	Total         string `json:"total"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdateAt      string `json:"update_at"`
	UpdatedBy     string `json:"updated_by"`
}

type ReqInquiryPI struct {
	IdDivisi   string    `json:"id_divisi"`
	RumahSakit string    `json:"rumah_sakit"`
	Alamat     string    `json:"alamat"`
	JatuhTempo string    `json:"jatuh_tempo"`
	Item       []ReqItem `json:"item"`
}

type ReqItem struct {
	Kat         string `json:"kat"`
	NamaBarang  string `json:"nama_barang"`
	Quantity    string `json:"quantity"`
	HargaSatuan string `json:"harga_satuan"`
	Discount    string `json:"discount"`
}

type ResInquiryPI struct {
	IdDivisi     string    `json:"id_divisi"`
	RumahSakit   string    `json:"rumah_sakit"`
	Alamat       string    `json:"alamat"`
	NomorInvoice string    `json:"nomor_invoice"`
	NomorPO      string    `json:"nomor_po"`
	NomorSI      string    `json:"nomor_si"`
	Tanggal      string    `json:"tanggal"`
	JatuhTempo   string    `json:"jatuh_tempo"`
	SubTotal     string    `json:"sub_total"`
	PajakPPN     string    `json:"pajak_ppn"`
	Total        string    `json:"total"`
	SubTotalRP   string    `json:"RP_sub_total"`
	PajakPPNRP   string    `json:"RP_pajak_ppn"`
	TotalRP      string    `json:"RP_total"`
	Item         []ResItem `json:"item"`
}

type ResItem struct {
	Kat            string `json:"kat"`
	NamaBarang     string `json:"nama_barang"`
	Quantity       string `json:"quantity"`
	HargaSatuan    string `json:"harga_satuan"`
	Discount       string `json:"discount"`
	SubTotalItem   string `json:"sub_total_item"`
	SubTotalItemRP string `json:"RP_sub_total_item"`
}

type Customer struct {
	Name               string `json:"name"`
	AddressCompany     string `json:"address_company"`
	NPWPAddress        string `json:"npwp_address"`
	NPWP               string `json:"npwp"`
	IpakNumber         string `json:"ipak_number"`
	FactureAddress     string `json:"facture_address"`
	CityFacture        string `json:"city_facture"`
	ZipCodeFacture     string `json:"zip_code_facture"`
	NumberPhoneFacture string `json:"number_phone_facture"`
	EmailFacture       string `json:"email_facture"`
	FaxFacture         string `json:"fax_facture"`
	PicFacture         string `json:"pic_facture"`
	ItemAddress        string `json:"item_address"`
	CityItem           string `json:"city_item"`
	ZipCodeItem        string `json:"zip_code_item"`
	NumberPhoneItem    string `json:"number_phone_item"`
	EmailItem          string `json:"email_item"`
	FaxItem            string `json:"fax_item"`
	PicItem            string `json:"pic_item"`
	ContactPerson      string `json:"contact_person"`
	TaxCodeID          string `json:"tax_code_id"`
	Top                string `json:"top"`
}
