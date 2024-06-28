package model

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
	ID     string `json:"id"`
	Divisi string `json:"divisi"`
}

type PerformanceInvoice struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	Status        string `json:"status"`
	Divisi        string `json:"divisi"`
	InvoiceNumber string `json:"invoice_number"`
	PONumber      string `json:"po_number"`
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
	Handphone          string `json:"handphone"`
}

type PerformanceInvoiceDetail struct {
	ID              int             `json:"id,omitempty"`
	CustomerID      int             `json:"customer_id,omitempty"`
	Status          string          `json:"status,omitempty"`
	Divisi          string          `json:"divisi,omitempty"`
	InvoiceNumber   string          `json:"invoice_number,omitempty"`
	PONumber        string          `json:"po_number,omitempty"`
	DueDate         string          `json:"due_date,omitempty"`
	DoctorName      string          `json:"doctor_name,omitempty"`
	PatientName     string          `json:"patient_name,omitempty"`
	TanggalTindakan string          `json:"tanggal_tindakan,omitempty"`
	RM              string          `json:"rm,omitempty"`
	NumberSI        string          `json:"number_si,omitempty"`
	SubTotal        string          `json:"sub_total,omitempty"`
	SubTotalRP      string          `json:"RP_sub_total"`
	Pajak           string          `json:"pajak,omitempty"`
	PajakPPNRP      string          `json:"RP_pajak_ppn"`
	Total           string          `json:"total,omitempty"`
	TotalRP         string          `json:"RP_total"`
	Reason          string          `json:"reason,omitempty"`
	ItemDetailPI    []ResItemDetail `json:"item_detail_pi,omitempty"`
}

type ResItemDetail struct {
	Id             int    `json:"id,omitempty"`
	Kat            string `json:"kat,omitempty"`
	NamaBarang     string `json:"nama_barang,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	HargaSatuan    string `json:"harga_satuan,omitempty"`
	Discount       string `json:"discount,omitempty"`
	SubTotalItem   string `json:"sub_total_item,omitempty"`
	RPSubTotalItem string `json:"rp_sub_total_item,omitempty"`
}
