package proformaInvoice

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func Posting(c *gin.Context) {
	// Ini untuk posting

	var input model.ResInquiryPI
	var newID int

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	log.Println("[--->]", "Data Input :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	var idCustomer int

	QueryCekRS := `
		select COALESCE(id, 0) AS id from customer where nama_company = $1
	`
	err = tx.QueryRow(ctx, QueryCekRS, input.RumahSakit).Scan(&idCustomer)

	if err != nil {
		log.Println("[--->]", "Error Get id Perusahaan :", err)
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if idCustomer == 0 {
		utility.ResponseError(c, "Perusahaan Tidak Ditemukan")
		return
	} else {
		log.Println("Perusahaan Tersedia, Proses di lanjut !")
	}

	if input.IdDivisi == "Ortopedi" {
		log.Println("[--->]", "Proses Input Ortopedi!")
		QeuryInputPI := `
			INSERT INTO performance_invoice (divisi, customer_id, invoice_number, number_si, tanggal_tindakan, due_date, total, sub_total, pajak, created_at , created_by , update_at , updated_by, doctor_name , patient_name , status )
			VALUES 
			    (
			        
			     $1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), 'SALES', NOW(), 'SALES', $10 , $11 , 'Diproses'
			        
			     )
			RETURNING id
		`

		err = tx.QueryRow(ctx, QeuryInputPI, input.IdDivisi, idCustomer, input.NomorInvoice, input.NomorSI, input.TanggalTindakan, input.JatuhTempo, input.TotalRP, input.SubTotalRP, input.PajakPPNRP, input.NamaDokter, input.NamaPasien).Scan(&newID)

		if err != nil {
			log.Println("[--->]", "Error Query Input PI Ortopedi:", err)
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}

		if len(input.Item) > 0 {
			for _, item := range input.Item {
				QueryItem := `
					insert into order_items (
						pi_id, "name", quantity, price, discount, 
						sub_total, kat, 
						created_at, created_by, update_at, updated_by
						) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), 'sales', NOW(), 'sales')
				`

				_, err = tx.Exec(ctx, QueryItem,
					newID, item.NamaBarang, item.Quantity, item.HargaSatuan,
					item.Discount, item.SubTotalItemRP, item.Kat)

				if err != nil {
					log.Println("[--->]", "Error Query Input Item Order :", err)
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}

				//  Pengaruhi Jumlah Stock Barang

				QueryStock := `
					UPDATE stock_items
					SET total = total::integer - $1
					WHERE name = $2;
				`

				_, err = tx.Exec(ctx, QueryStock,
					item.Quantity, item.NamaBarang)

				if err != nil {
					log.Println("[--->]", "Error Query Update Stock Barang :", err)
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}
			}
		}
	} else if input.IdDivisi == "Radiologi" {
		log.Println("[--->]", "Proses Input Radiologi!")
		QeuryInputPI := `
			INSERT INTO performance_invoice (divisi, customer_id, invoice_number, number_si, due_date, total, sub_total, pajak, created_at , created_by , update_at , updated_by, status  )
			VALUES 
			    (
			        
			     $1, $2, $3, $4, $5, $6, $7, $8, NOW(), 'SALES', NOW(), 'SALES', 'Diproses'
			        
			     )
			RETURNING id
		`

		err = tx.QueryRow(ctx, QeuryInputPI, input.IdDivisi, idCustomer, input.NomorInvoice, input.NomorSI, input.JatuhTempo, input.TotalRP, input.SubTotalRP, input.PajakPPNRP).Scan(&newID)

		if err != nil {
			log.Println("[--->]", "Error Query Input PI Ortopedi:", err)
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}

		if len(input.Item) > 0 {
			for _, item := range input.Item {
				QueryItem := `
					insert into order_items (
						pi_id, "name", quantity, price, discount, 
						sub_total, kat, 
						created_at, created_by, update_at, updated_by
						) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), 'sales', NOW(), 'sales')
				`

				_, err = tx.Exec(ctx, QueryItem,
					newID, item.NamaBarang, item.Quantity, item.HargaSatuan,
					item.Discount, item.SubTotalItemRP, item.Kat)

				if err != nil {
					log.Println("[--->]", "Error Query Input Item Order :", err)
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}

				QueryStock := `
					UPDATE stock_items
					SET total = total::integer - $1
					WHERE name = $2;
				`

				_, err = tx.Exec(ctx, QueryStock,
					item.Quantity, item.NamaBarang)

				if err != nil {
					log.Println("[--->]", "Error Query Update Stock Barang :", err)
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}
			}
		}
	}

	//  Ketika Berhasil Insert Ke Teble Piutang

	QueryPiutang := `
					insert into piutang (
						nama, nominal, amount, tanggal
						) 
					VALUES ($1, $2, $3, NOW())
				`

	_, err = tx.Exec(ctx, QueryPiutang,
		input.RumahSakit, input.Total, input.TotalRP)

	if err != nil {
		log.Println("[--->]", "Error Query Input Piutang :", err)
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	//  Ketika Berhasil Insert Ke Table Pemasukan

	QueryPemasukan := `
					insert into pemasukan (
						nama, nominal, amount, tanggal
						) 
					VALUES ($1, $2, $3, NOW())
				`

	_, err = tx.Exec(ctx, QueryPemasukan,
		input.RumahSakit, input.Total, input.TotalRP)

	if err != nil {
		log.Println("[--->]", "Error Query Input Pemasukan :", err)
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data added successfully", "status": true})

}
