package service

import "time"

type StockBarang struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Total     string    `json:"total,omitempty"`
	Price     string    `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CreatedBy string    `json:"created_by,omitempty"`
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
	IdDivisi        string    `json:"id_divisi"`
	RumahSakit      string    `json:"rumah_sakit"`
	Alamat          string    `json:"alamat"`
	JatuhTempo      string    `json:"jatuh_tempo"`
	NamaDokter      string    `json:"nama_dokter"`
	NamaPasien      string    `json:"nama_pasien"`
	IdRumahSakit    string    `json:"id_rumah_sakit"`
	TanggalTindakan string    `json:"tanggal_tindakan"`
	RM              string    `json:"rm"`
	Item            []ReqItem `json:"item"`
}

type ReqItem struct {
	Kat         string `json:"kat"`
	NamaBarang  string `json:"nama_barang"`
	Quantity    string `json:"quantity"`
	HargaSatuan string `json:"harga_satuan"`
	Discount    string `json:"discount"`
}

type ResInquiryPI struct {
	IdDivisi        string    `json:"id_divisi"`
	IdRumahSakit    string    `json:"id_rumah_sakit"`
	RumahSakit      string    `json:"rumah_sakit"`
	Alamat          string    `json:"alamat"`
	NomorInvoice    string    `json:"nomor_invoice"`
	NomorPO         string    `json:"nomor_po"`
	NomorSI         string    `json:"nomor_si"`
	Tanggal         string    `json:"tanggal"`
	JatuhTempo      string    `json:"jatuh_tempo"`
	SubTotal        string    `json:"sub_total"`
	PajakPPN        string    `json:"pajak_ppn"`
	Total           string    `json:"total"`
	SubTotalRP      string    `json:"RP_sub_total"`
	PajakPPNRP      string    `json:"RP_pajak_ppn"`
	TotalRP         string    `json:"RP_total"`
	NamaDokter      string    `json:"nama_dokter,omitempty"`
	NamaPasien      string    `json:"nama_pasien,omitempty"`
	TanggalTindakan string    `json:"tanggal_tindakan,omitempty"`
	RM              string    `json:"rm,omitempty"`
	Item            []ResItem `json:"item"`
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
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	AddressCompany     string `json:"address_company,omitempty"`
	NPWPAddress        string `json:"npwp_address,omitempty"`
	NPWP               string `json:"npwp,omitempty"`
	IpakNumber         string `json:"ipak_number,omitempty"`
	FactureAddress     string `json:"facture_address,omitempty"`
	CityFacture        string `json:"city_facture,omitempty"`
	ZipCodeFacture     string `json:"zip_code_facture,omitempty"`
	NumberPhoneFacture string `json:"number_phone_facture,omitempty"`
	EmailFacture       string `json:"email_facture,omitempty"`
	FaxFacture         string `json:"fax_facture,omitempty"`
	PicFacture         string `json:"pic_facture,omitempty"`
	ItemAddress        string `json:"item_address,omitempty"`
	CityItem           string `json:"city_item,omitempty"`
	ZipCodeItem        string `json:"zip_code_item,omitempty"`
	NumberPhoneItem    string `json:"number_phone_item,omitempty"`
	EmailItem          string `json:"email_item,omitempty"`
	FaxItem            string `json:"fax_item,omitempty"`
	PicItem            string `json:"pic_item,omitempty"`
	ContactPerson      string `json:"contact_person,omitempty"`
	TaxCodeID          string `json:"tax_code_id,omitempty"`
	Top                string `json:"top,omitempty"`
}
