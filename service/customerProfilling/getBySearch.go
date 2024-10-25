package customerProfilling

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"time"
)

func GetBySearch(c *gin.Context) {

	var input model.Customer

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

	log.Println("[--->]", "Data Input :", input)

	query := `
		SELECT 
			id,
			COALESCE(nama_perusahaan, '') AS nama_company,
			COALESCE(address_perusahaan, '') AS address_company,
			COALESCE(npwp_address_perusahaan , '') AS npwp_address,
			COALESCE(npwp_perusahaan , '') AS npwp,
			COALESCE(ipak_number_perusahaan , '') AS ipak_number,
			COALESCE(alamat_pengirim_facture_perusahaan , '') AS facture_address,
			COALESCE(kota_perusahaan , '') AS city_facture,
			COALESCE(kode_pos_perusahaan , '') AS zip_code_facture,
			COALESCE(telpon_perusahaan , '') AS number_phone_facture,
			COALESCE(email_perusahaan , '') AS email_facture,
			COALESCE(pic_perusahaan , '') AS pic_facture,
			COALESCE(alamat_pengirim_dokter , '') AS item_address,
			COALESCE(kota_dokter , '') AS city_item,
			COALESCE(kode_pos_dokter , '') AS zip_code_item,
			COALESCE(telpon_dokter , '') AS number_phone_item,
			COALESCE(email_dokter , '') AS email_item,
			COALESCE(pic_dokter , '') AS pic_item,
			COALESCE(cp_dokter , '') AS contact_person,
			COALESCE(CAST(tax_code  AS TEXT), '0') AS tax_code_id,
			COALESCE(term_of_payment , '') AS top,
			coalesce(nama_dokter, '')as nama_dokter ,
			kategori_divisi 
		FROM 
			public.customer
		where nama_perusahaan = $1;
	`

	rows, err := tx.Query(ctx, query, input.Name)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}
	defer rows.Close()

	var response []model.Customer

	for rows.Next() {
		var tc model.Customer
		if err := rows.Scan(
			&tc.ID,
			&tc.Name,
			&tc.AddressCompany,
			&tc.NPWPAddress,
			&tc.NPWP,
			&tc.IpakNumber,
			&tc.FactureAddress,
			&tc.CityFacture,
			&tc.ZipCodeFacture,
			&tc.NumberPhoneFacture,
			&tc.EmailFacture,
			&tc.PicFacture,
			&tc.ItemAddress,
			&tc.CityItem,
			&tc.ZipCodeItem,
			&tc.NumberPhoneItem,
			&tc.EmailItem,
			&tc.PicItem,
			&tc.ContactPerson,
			&tc.TaxCodeID,
			&tc.Top,
			&tc.DocktorName,
			&tc.KategoriDivisi,
		); err != nil {
			log.Println("[--->]", "Error Scan Data :", error(err))
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Data", "status": false})
			return
		}
		response = append(response, tc)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over tax code rows", "status": false})
		return
	}

	var jumlahPI int

	// Hitung Jumlah Order PI
	for _, customer := range response {
		log.Println("Customer : ", customer)

		hitungPI := `
		SELECT COUNT(*) AS total_invoices
		FROM performance_invoice
		WHERE customer = $1;
		`
		var hitungpi int
		err := tx.QueryRow(ctx, hitungPI, input.Name).Scan(&hitungpi)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}

		jumlahPI = jumlahPI + hitungpi
	}

	if len(response) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Data Ditemukan !",
			"jumlah_pi": jumlahPI,
			"customer":  response,
			//"riwayat_order": totalQuantity,
			"status": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "Data": []model.Customer{}, "status": true})
	}

}

func ListProfile(c *gin.Context) {

	var input model.RequestID

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

	queryList := `
		select 
		    id,
			COALESCE(nama_perusahaan, '' )as nama_perusahaan, 
			COALESCE(telpon_perusahaan, '' )as telpon_perusahaan, 
			COALESCE(nama_dokter, '' )as nama_dokter, 
			COALESCE(telpon_dokter, '' )as telpon_dokter 
		from customer pl where nama_perusahaan  = $1
	`

	row, err := tx.Query(ctx, queryList, input.Nama)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	defer row.Close()

	var Data []model.Profile

	for row.Next() {
		var ambil model.Profile
		err := row.Scan(
			&ambil.ID,
			&ambil.NamaPerusahaan,
			&ambil.NotelpPerusahaan,
			&ambil.NamaDokter,
			&ambil.NotelpDokter,
		)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Data = append(Data, ambil)
	}

	if err := row.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Data) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Data, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}
}

func DetailProfile(c *gin.Context) {

	var input model.RequestID

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

	query := `
		SELECT 
			COALESCE(c.id, 0) AS id,
			COALESCE(c.nama_perusahaan, '') AS nama_perusahaan,
			COALESCE(c.address_perusahaan, '') AS address_perusahaan,
			COALESCE(c.npwp_address_perusahaan, '') AS npwp_address_perusahaan,
			COALESCE(c.npwp_perusahaan, '') AS npwp_perusahaan,
			COALESCE(c.ipak_number_perusahaan, '') AS ipak_number_perusahaan,
			COALESCE(c.alamat_pengirim_facture_perusahaan, '') AS alamat_pengirim_facture_perusahaan,
			COALESCE(c.kota_perusahaan, '') AS kota_perusahaan,
			COALESCE(c.kode_pos_perusahaan, '') AS kode_pos_perusahaan,
			COALESCE(c.telpon_perusahaan, '') AS telpon_perusahaan,
			COALESCE(c.email_perusahaan, '') AS email_perusahaan,
			COALESCE(c.nama_dokter, '') AS nama_dokter,
			COALESCE(c.alamat_pengirim_dokter, '') AS alamat_pengirim_dokter,
			COALESCE(c.npwp_dokter, '') AS npwp_dokter,
			COALESCE(c.telpon_dokter, '') AS telpon_dokter,
			COALESCE(c.email_dokter, '') AS email_dokter,
			COALESCE(c.pic_dokter, '') AS pic_dokter,
			COALESCE(c.kota_dokter, '') AS kota_dokter,
			COALESCE(c.kode_pos_dokter, '') AS kode_pos_dokter,
			COALESCE(c.handphone_dokter, '') AS handphone_dokter,
			COALESCE(c.kode_pajak_dokter, '') AS kode_pajak_dokter,
			COALESCE(c.tax_code, 0) AS tax_code,
			COALESCE(c.verifikasi_dokter, '') AS verifikasi_dokter,
			COALESCE(TO_CHAR(c.created_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS created_at,
			COALESCE(c.created_by, '') AS created_by,
			COALESCE(TO_CHAR(c.updated_at, 'YYYY-MM-DD HH24:MI:SS'), '') AS updated_at,
			COALESCE(c.updated_by, '') AS updated_by,
			COALESCE(c.pembuat_cp_dokter, '') AS pembuat_cp_dokter,
			COALESCE(c.term_of_payment, '') AS term_of_payment,
			COALESCE(c.kategori_divisi, '') AS kategori_divisi,
			COALESCE(c.cp_dokter, '') AS cp_dokter,
			COALESCE(c.pic_perusahaan, '') AS pic_perusahaan
		FROM customer c where c.id = $1
	`

	// Menjalankan query dan mengisi ke dalam model
	var customer model.ProfileDetail
	err = tx.QueryRow(context.Background(), query, input.ID).Scan(
		&customer.ID,
		&customer.NamaPerusahaan,
		&customer.AddressPerusahaan,
		&customer.NPWPAddressPerusahaan,
		&customer.NPWPPerusahaan,
		&customer.IPAKNumberPerusahaan,
		&customer.AlamatPengirimFacturePerusahaan,
		&customer.KotaPerusahaan,
		&customer.KodePosPerusahaan,
		&customer.TelponPerusahaan,
		&customer.EmailPerusahaan,
		&customer.NamaDokter,
		&customer.AlamatPengirimDokter,
		&customer.NPWPDokter,
		&customer.TelponDokter,
		&customer.EmailDokter,
		&customer.PicDokter,
		&customer.KotaDokter,
		&customer.KodePosDokter,
		&customer.HandphoneDokter,
		&customer.KodePajakDokter,
		&customer.TaxCode,
		&customer.VerifikasiDokter,
		&customer.CreatedAt,
		&customer.CreatedBy,
		&customer.UpdatedAt,
		&customer.UpdatedBy,
		&customer.PembuatCPDokter,
		&customer.TermOfPayment,
		&customer.KategoriDivisi,
		&customer.CPDokter,
		&customer.PicPerusahaan,
	)

	if err != nil {
		log.Println("Error querying customer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": customer, "status": true})

}

func EditDetailProfile(c *gin.Context) {

	var input model.CustomerNew

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {

			log.Println("Insert failed:", err)
			c.JSON(400, gin.H{"message": err})
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {

			log.Println("Insert failed:", err)
			c.JSON(400, gin.H{"message": err})
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	input.NamaDokter = utility.ToUpperCase(input.NamaDokter)
	input.NamaPerusahaan = utility.ToUpperCase(input.NamaPerusahaan)

	query := `
			UPDATE customer SET
			    nama_perusahaan = $1,
				address_perusahaan = $2,
				npwp_address_perusahaan = $3,
				npwp_perusahaan = $4,
				ipak_number_perusahaan = $5,
				alamat_pengirim_facture_perusahaan = $6,
				kota_perusahaan = $7,
				kode_pos_perusahaan = $8,
				telpon_perusahaan = $9,
				email_perusahaan = $10,
				nama_dokter = $11,
				alamat_pengirim_dokter = $12,
				npwp_dokter = $13,
				telpon_dokter = $14,
				email_dokter = $15,
				pic_dokter = $16,
				kota_dokter = $17,
				kode_pos_dokter = $18,
				handphone_dokter = $19,
				kode_pajak_dokter = $20,
				cp_dokter = $21,
				verifikasi_dokter = $22,
				created_at = $23,
				created_by = $24,
				updated_at = $25,
				updated_by = $26,
				pembuat_cp_dokter = $27,
				term_of_payment = $28,
				kategori_divisi = $29
			WHERE id = $30

			`

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
		input.ID,
	)

	if err != nil {
		tx.Rollback(ctx)
		log.Println("Insert failed:", err)
		c.JSON(400, gin.H{"message": err})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(400, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully", "status": true})

}
