package proformaInvoice

import (
	"backend_project_fismed/model"
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

	if input.IdDivisi == "1" {
		log.Println("[--->]", "Proses Input Ortopedi!")
		QeuryInputPI := `
			insert into performance_invoice (
			   customer_id, sub_total, 
			   status, divisi, invoice_number, po_number, due_date, 
			   created_at, created_by, update_at, updated_by, total, pajak, number_si
			) VALUES (
			$1, $2, 'Diproses', 'Ortopedi', $3, $4, $5, NOW(), 'sales', NOW(), 'sales', $6, $7, $8 
			)
			RETURNING id
		`

		err = tx.QueryRow(ctx, QeuryInputPI, input.IdRumahSakit, input.SubTotalRP, input.NomorInvoice,
			input.NomorPO, input.JatuhTempo, input.TotalRP, input.PajakPPNRP, input.NomorSI).Scan(&newID)

		if err != nil {
			log.Println("[--->]", "Error Query Input PI :", err)
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
			}
		}
	} else if input.IdDivisi == "2" {
		log.Println("[--->]", "Proses Input Radiologi!")
		QeuryInputPI := `
			insert into performance_invoice (
			   	customer_id, sub_total, 
			   	status, divisi, invoice_number, po_number, due_date, 
			   	created_at, created_by, update_at, updated_by, total, pajak, 
			   	doctor_name, patient_name, number_si,
			   	tanggal_tindakan, rm
			) VALUES (
			$1, $2, 'Diproses', 'Ortopedi', $3, $4, $5, NOW(), 'sales', NOW(), 'sales', $6, $7, $8, $9, $10, $11, $12  
			)
			RETURNING id
		`

		err = tx.QueryRow(ctx, QeuryInputPI, input.IdRumahSakit, input.SubTotalRP, input.NomorInvoice,
			input.NomorPO, input.JatuhTempo, input.TotalRP, input.PajakPPNRP,
			input.NamaDokter, input.NamaPasien, input.NomorSI, input.TanggalTindakan, input.RM).Scan(&newID)

		if err != nil {
			log.Println("[--->]", "Error Query Input PI :", err)
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
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data added successfully", "status": true})

}
