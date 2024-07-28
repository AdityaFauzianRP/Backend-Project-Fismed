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
	ID     string `json:"id,omitempty"`
	Divisi string `json:"divisi,omitempty"`
	Export string `json:"export,omitempty"`
}

type PerformanceInvoice struct {
	ID            int    `json:"id,omitempty"`
	CustomerID    int    `json:"customer_id,omitempty"`
	Status        string `json:"status,omitempty"`
	Divisi        string `json:"divisi,omitempty"`
	InvoiceNumber string `json:"invoice_number,omitempty"`
	PONumber      string `json:"po_number,omitempty"`
	SubTotal      string `json:"sub_total,omitempty"`
	Pajak         string `json:"pajak,omitempty"`
	Total         string `json:"total,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatedBy     string `json:"created_by,omitempty"`
	UpdateAt      string `json:"update_at,omitempty"`
	UpdatedBy     string `json:"updated_by,omitempty"`
	NamaCompany   string `json:"nama_company,omitempty"`
}

type ReqInquiryPI struct {
	IdDivisi        string    `json:"id_divisi,omitempty"`
	RumahSakit      string    `json:"rumah_sakit,omitempty"`
	Alamat          string    `json:"alamat,omitempty"`
	JatuhTempo      string    `json:"jatuh_tempo,omitempty"`
	NamaDokter      string    `json:"nama_dokter,omitempty"`
	NamaPasien      string    `json:"nama_pasien,omitempty"`
	IdRumahSakit    string    `json:"id_rumah_sakit,omitempty"`
	TanggalTindakan string    `json:"tanggal_tindakan,omitempty"`
	RM              string    `json:"rm,omitempty"`
	Item            []ReqItem `json:"item,omitempty"`
}

type ReqItem struct {
	Kat         string `json:"kat,omitempty"`
	NamaBarang  string `json:"nama_barang,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	HargaSatuan string `json:"harga_satuan,omitempty"`
	Discount    string `json:"discount,omitempty"`
}

type ResInquiryPI struct {
	IdDivisi        string    `json:"id_divisi,omitempty"`
	IdRumahSakit    string    `json:"id_rumah_sakit,omitempty"`
	RumahSakit      string    `json:"rumah_sakit,omitempty"`
	Alamat          string    `json:"alamat,omitempty"`
	NomorInvoice    string    `json:"nomor_invoice,omitempty"`
	NomorPO         string    `json:"nomor_po,omitempty"`
	NomorSI         string    `json:"nomor_si,omitempty"`
	Tanggal         string    `json:"tanggal,omitempty"`
	JatuhTempo      string    `json:"jatuh_tempo,omitempty"`
	SubTotal        string    `json:"sub_total,omitempty"`
	PajakPPN        string    `json:"pajak_ppn,omitempty"`
	Total           string    `json:"total,omitempty"`
	SubTotalRP      string    `json:"RP_sub_total,omitempty"`
	PajakPPNRP      string    `json:"RP_pajak_ppn,omitempty"`
	TotalRP         string    `json:"RP_total,omitempty"`
	NamaDokter      string    `json:"nama_dokter,omitempty"`
	NamaPasien      string    `json:"nama_pasien,omitempty"`
	TanggalTindakan string    `json:"tanggal_tindakan,omitempty"`
	RM              string    `json:"rm,omitempty"`
	Item            []ResItem `json:"item,omitempty"`
}

type ResItem struct {
	Kat            string `json:"kat,omitempty"`
	NamaBarang     string `json:"nama_barang,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	HargaSatuan    string `json:"harga_satuan,omitempty"`
	Discount       string `json:"discount,omitempty"`
	SubTotalItem   string `json:"sub_total_item,omitempty"`
	SubTotalItemRP string `json:"RP_sub_total_item,omitempty"`
}

type Customer struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	NameCompany        string `json:"nama_company,omitempty"`
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
	Handphone          string `json:"handphone,omitempty"`
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
	SubTotalRP      string          `json:"RP_sub_total,omitempty"`
	Pajak           string          `json:"pajak,omitempty"`
	PajakPPNRP      string          `json:"RP_pajak_ppn,omitempty"`
	Total           string          `json:"total,omitempty"`
	TotalRP         string          `json:"RP_total,omitempty"`
	Reason          string          `json:"reason,omitempty"`
	Tanggal         string          `json:"tanggal,omitempty"`
	Customer        string          `json:"nama_customer,omitempty"`
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
