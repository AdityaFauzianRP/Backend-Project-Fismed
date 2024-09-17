package model

import "time"

type CustomerNew struct {
	ID                              int       `json:"id"`
	NamaPerusahaan                  string    `json:"nama_perusahaan"`
	AddressPerusahaan               string    `json:"address_perusahaan"`
	NPWPAddressPerusahaan           string    `json:"npwp_address_perusahaan"`
	NPWPPerusahaan                  string    `json:"npwp_perusahaan"`
	IPAKNumberPerusahaan            string    `json:"ipak_number_perusahaan"`
	AlamatPengirimFacturePerusahaan string    `json:"alamat_pengirim_facture_perusahaan"`
	KotaPerusahaan                  string    `json:"kota_perusahaan"`
	KodePosPerusahaan               string    `json:"kode_pos_perusahaan"`
	TelponPerusahaan                string    `json:"telpon_perusahaan"`
	EmailPerusahaan                 string    `json:"email_perusahaan"`
	NamaDokter                      string    `json:"nama_dokter"`
	AlamatPengirimDokter            string    `json:"alamat_pengirim_dokter"`
	NPWPDokter                      string    `json:"npwp_dokter"`
	TelponDokter                    string    `json:"telpon_dokter"`
	EmailDokter                     string    `json:"email_dokter"`
	PICDokter                       string    `json:"pic_dokter"`
	KotaDokter                      string    `json:"kota_dokter"`
	KodePosDokter                   string    `json:"kode_pos_dokter"`
	HandphoneDokter                 string    `json:"handphone_dokter"`
	KodePajakDokter                 string    `json:"kode_pajak_dokter"`
	CPDokter                        string    `json:"cp_dokter"`
	VerifikasiDokter                string    `json:"verifikasi_dokter"`
	CreatedAt                       time.Time `json:"created_at"`
	CreatedBy                       string    `json:"created_by"`
	UpdatedAt                       time.Time `json:"updated_at"`
	UpdatedBy                       string    `json:"updated_by"`
	PembuatCPDokter                 string    `json:"pembuat_cp_dokter"`
	TermOfPayment                   string    `json:"term_of_payment"`
	KategoriDivisi                  string    `json:"kategori_divisi"`
}
