package customerProfilling

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
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
			coalesce(nama_dokter, '')as nama_dokter 
		FROM 
			public.customer
		where nama_perusahaan = $1;
	`

	rows, err := tx.Query(ctx, query, input.Name)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
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
		hitungPI := `
		SELECT COUNT(*) AS total_invoices
		FROM performance_invoice
		WHERE customer_id = $1;
		`
		var hitungpi int
		err := tx.QueryRow(ctx, hitungPI, customer.ID).Scan(&hitungpi)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
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
