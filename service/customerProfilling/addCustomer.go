package customerProfilling

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"time"
)

func Add(c *gin.Context) {

	var input model.CustomerNew

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	input.NamaDokter = utility.ToUpperCase(input.NamaDokter, "0")
	input.NamaPerusahaan = utility.ToUpperCase(input.NamaPerusahaan, input.KategoriDivisi)

	if input.KategoriDivisi == constanta.CustomerRS {

		log.Println("Customer Rumah Sakit !")

		query := `
			INSERT INTO customer (
				nama_perusahaan,
				address_perusahaan,
				npwp_address_perusahaan,
				npwp_perusahaan,
				ipak_number_perusahaan,
				alamat_pengirim_facture_perusahaan,
				kota_perusahaan,
				kode_pos_perusahaan,
				telpon_perusahaan,
				email_perusahaan,
				nama_dokter,
				alamat_pengirim_dokter,
				npwp_dokter,
				telpon_dokter,
				email_dokter,
				pic_dokter,
				kota_dokter,
				kode_pos_dokter,
				handphone_dokter,
				kode_pajak_dokter,
				cp_dokter,
				verifikasi_dokter,
				created_at,
				created_by,
				updated_at,
				updated_by,
				pembuat_cp_dokter,
				term_of_payment,
				kategori_divisi
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29
			)`

		// Data yang akan diinsert
		_, err = tx.Exec(ctx, query,
			input.NamaPerusahaan,
			input.AddressPerusahaan,
			input.NPWPAddressPerusahaan,
			input.NPWPPerusahaan,
			input.IPAKNumberPerusahaan,
			input.AlamatPengirimFacturePerusahaan,
			input.KotaPerusahaan,
			input.KodePosPerusahaan,
			input.TelponPerusahaan,
			input.EmailPerusahaan,
			input.NamaDokter,
			input.AlamatPengirimDokter,
			input.NPWPDokter,
			input.TelponDokter,
			input.EmailDokter,
			input.PICDokter,
			input.KotaDokter,
			input.KodePosDokter,
			input.HandphoneDokter,
			input.KodePajakDokter,
			input.CPDokter,
			input.VerifikasiDokter,
			time.Now(),
			"admin",
			time.Now(),
			"admin",
			input.PembuatCPDokter,
			input.TermOfPayment,
			input.KategoriDivisi,
		)

	} else if input.KategoriDivisi == constanta.CustomerNonRS {

		log.Println("Customer Non Rumah Sakit ! ")

		query := `
			INSERT INTO customer (
				nama_perusahaan,
				address_perusahaan,
				npwp_address_perusahaan,
				npwp_perusahaan,
				ipak_number_perusahaan,
				alamat_pengirim_facture_perusahaan,
				kota_perusahaan,
				kode_pos_perusahaan,
				telpon_perusahaan,
				email_perusahaan,
				nama_dokter,
				alamat_pengirim_dokter,
				npwp_dokter,
				telpon_dokter,
				email_dokter,
				pic_dokter,
				kota_dokter,
				kode_pos_dokter,
				handphone_dokter,
				kode_pajak_dokter,
				cp_dokter,
				verifikasi_dokter,
				created_at,
				created_by,
				updated_at,
				updated_by,
				pembuat_cp_dokter,
				term_of_payment,
				kategori_divisi
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29
			)`

		// Data yang akan diinsert
		_, err = tx.Exec(ctx, query,
			input.NamaPerusahaan,
			input.AddressPerusahaan,
			input.NPWPAddressPerusahaan,
			input.NPWPPerusahaan,
			input.IPAKNumberPerusahaan,
			input.AlamatPengirimFacturePerusahaan,
			input.KotaPerusahaan,
			input.KodePosPerusahaan,
			input.TelponPerusahaan,
			input.EmailPerusahaan,
			input.NamaDokter,
			input.AlamatPengirimDokter,
			input.NPWPDokter,
			input.TelponDokter,
			input.EmailDokter,
			input.PICDokter,
			input.KotaDokter,
			input.KodePosDokter,
			input.HandphoneDokter,
			input.KodePajakDokter,
			input.CPDokter,
			input.VerifikasiDokter,
			time.Now(),
			"admin",
			time.Now(),
			"admin",
			input.PembuatCPDokter,
			input.TermOfPayment,
			input.KategoriDivisi,
		)

	} else if input.KategoriDivisi == constanta.CustomerAsSupplier {

		log.Println("Customer As Supplier ! ")

		query := `
			INSERT INTO customer (
				nama_perusahaan,
				address_perusahaan,
				npwp_address_perusahaan,
				npwp_perusahaan,
				ipak_number_perusahaan,
				alamat_pengirim_facture_perusahaan,
				kota_perusahaan,
				kode_pos_perusahaan,
				telpon_perusahaan,
				email_perusahaan,
				nama_dokter,
				alamat_pengirim_dokter,
				npwp_dokter,
				telpon_dokter,
				email_dokter,
				pic_dokter,
				kota_dokter,
				kode_pos_dokter,
				handphone_dokter,
				kode_pajak_dokter,
				cp_dokter,
				verifikasi_dokter,
				created_at,
				created_by,
				updated_at,
				updated_by,
				pembuat_cp_dokter,
				term_of_payment,
				kategori_divisi
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29
			)`

		// Data yang akan diinsert
		_, err = tx.Exec(ctx, query,
			input.NamaPerusahaan,
			input.AddressPerusahaan,
			input.NPWPAddressPerusahaan,
			input.NPWPPerusahaan,
			input.IPAKNumberPerusahaan,
			input.AlamatPengirimFacturePerusahaan,
			input.KotaPerusahaan,
			input.KodePosPerusahaan,
			input.TelponPerusahaan,
			input.EmailPerusahaan,
			input.NamaDokter,
			input.AlamatPengirimDokter,
			input.NPWPDokter,
			input.TelponDokter,
			input.EmailDokter,
			input.PICDokter,
			input.KotaDokter,
			input.KodePosDokter,
			input.HandphoneDokter,
			input.KodePajakDokter,
			input.CPDokter,
			input.VerifikasiDokter,
			time.Now(),
			"admin",
			time.Now(),
			"admin",
			input.PembuatCPDokter,
			input.TermOfPayment,
			input.KategoriDivisi,
		)

	}

	if err != nil {
		tx.Rollback(ctx)
		log.Println("Insert failed:", err)
		c.JSON(400, gin.H{"message": err})
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(400, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully", "status": true})

}
