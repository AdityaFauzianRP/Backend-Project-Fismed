package proformaInvoice

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func EditPI(c *gin.Context) {
	//  Edit Data PI
	var input model.PerformanceInvoiceDetail
	var response model.PerformanceInvoiceDetail

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	// Proses Inquiry Edit

	response.ID = input.ID
	response.CustomerID = input.CustomerID
	response.Status = input.Status
	response.Divisi = input.Divisi
	response.InvoiceNumber = input.InvoiceNumber
	response.PONumber = input.PONumber
	response.DueDate = input.DueDate
	response.DoctorName = input.DoctorName
	response.PatientName = input.PatientName
	response.TanggalTindakan = input.TanggalTindakan
	response.RM = input.RM
	response.NumberSI = input.NumberSI

	var Subtotal int = 0
	response.ItemDetailPI = make([]model.ResItemDetail, len(input.ItemDetailPI))

	if len(input.ItemDetailPI) > 0 {

		for i, item := range input.ItemDetailPI {
			log.Println("Data Ke :", i)

			response.ItemDetailPI[i].Kat = item.Kat
			response.ItemDetailPI[i].NamaBarang = item.NamaBarang
			response.ItemDetailPI[i].Quantity = item.Quantity
			response.ItemDetailPI[i].HargaSatuan = item.HargaSatuan
			response.ItemDetailPI[i].Discount = item.Discount

			QuantitiInt, err := strconv.Atoi(item.Quantity)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Quantity format"})
				return
			}

			HargaSatuanInt, err := strconv.Atoi(item.HargaSatuan)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Harga Satuan format"})
				return
			}

			DiscountInt, err := strconv.Atoi(item.Discount)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Discount format"})
				return
			}

			subtotalperitem := QuantitiInt*HargaSatuanInt - (QuantitiInt * HargaSatuanInt * DiscountInt / 100)

			Subtotal = Subtotal + subtotalperitem

			subtotalperitemstring := strconv.Itoa(subtotalperitem)

			response.ItemDetailPI[i].SubTotalItem = subtotalperitemstring
			response.ItemDetailPI[i].RPSubTotalItem = "Rp. " + utility.FormatRupiah(subtotalperitemstring)
			response.ItemDetailPI[i].Id = item.Id
		}
	}

	log.Println("Subtotal :", Subtotal)

	pajak := Subtotal * 11 / 100
	total := pajak + Subtotal

	response.SubTotal = strconv.Itoa(Subtotal)
	response.Pajak = strconv.Itoa(pajak)
	response.Total = strconv.Itoa(total)
	response.TotalRP = "Rp. " + utility.FormatRupiah(response.Total)
	response.PajakPPNRP = "Rp. " + utility.FormatRupiah(response.Pajak)
	response.SubTotalRP = "Rp. " + utility.FormatRupiah(response.SubTotal)

	c.JSON(http.StatusOK, gin.H{"message": "Inquiry Performa Invoice Success !", "data": response, "status": true})
}

func PostingEdit_PI(c *gin.Context) {
	log.Println("Performa Invoice Posting")

	var input model.PerformanceInvoiceDetail

	//var response service.PerformanceInvoiceDetail

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	log.Println("Input Data :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	if input.Divisi == "Ortopedi" {
		query := `
			UPDATE performance_invoice
			SET
				customer_id = $1,
				sub_total = $2,
				status = $3,
				divisi = $4,
				invoice_number = $5,
				po_number = $6,
				due_date = $7,
				pajak = $8,
				total = $9,
				number_si = $10,
				update_at = now(),
				updated_by = 'sales'
			WHERE id = $11;
			`

		_, err = tx.Exec(context.Background(), query,
			input.CustomerID,
			input.SubTotalRP,
			input.Status,
			input.Divisi,
			input.InvoiceNumber,
			input.PONumber,
			input.DueDate,
			input.PajakPPNRP,
			input.TotalRP,
			input.NumberSI,
			input.ID,
		)

	} else if input.Divisi == "Radiologi" {
		query := `
			UPDATE performance_invoice
			SET
				customer_id = $1,
				sub_total = $2,
				status = $3,
				divisi = $4,
				invoice_number = $5,
				po_number = $6,
				due_date = $7,
				doctor_name = $8,
				patient_name = $9,
				pajak = $10,
				total = $11,
				tanggal_tindakan = $12,
				rm = $13,
				number_si = $14,
				update_at = now(),
				updated_by = 'sales'
			WHERE id = $15;
			`

		_, err = tx.Exec(context.Background(), query,
			input.CustomerID,
			input.SubTotalRP,
			input.Status,
			input.Divisi,
			input.InvoiceNumber,
			input.PONumber,
			input.DueDate,
			input.DoctorName,
			input.PatientName,
			input.PajakPPNRP,
			input.TotalRP,
			input.TanggalTindakan,
			input.RM,
			input.NumberSI,
			input.ID,
		)

	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if len(input.ItemDetailPI) > 0 {
		for _, detail := range input.ItemDetailPI {
			if detail.Id == 0 {
				QueryItem := `
					insert into order_items (
						pi_id, "name", quantity, price, discount, 
						sub_total, kat, 
						created_at, created_by, update_at, updated_by
						) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), 'sales', NOW(), 'sales')
				`

				_, err = tx.Exec(ctx, QueryItem,
					input.ID, detail.NamaBarang, detail.Quantity, detail.HargaSatuan,
					detail.Discount, detail.RPSubTotalItem, detail.Kat)

				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}
			} else {
				query := `
				UPDATE order_items
				SET
					name = $1,
					quantity = $2,
					price = $3,
					discount = $4,
					sub_total = $5,
					kat = $6,
					update_at = $7,
					updated_by = $8
				WHERE id = $9;
				`

				_, err = tx.Exec(context.Background(), query,
					detail.NamaBarang,
					detail.Quantity,
					detail.HargaSatuan,
					detail.Discount,
					detail.RPSubTotalItem,
					detail.Kat,
					time.Now(),
					"sales",
					input.ID,
				)

				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}
			}
		}
	}

	//  Ketika Berhasik Pengaruhi jumlah Stok

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Update Successfully", "status": true})

}

func EditAdmin(c *gin.Context) {
	var input model.PerformanceInvoiceDetail

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	log.Println("Input Data :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("tx.DBConnection :", tx)

	if input.Status == "Diterima" {
		log.Println("PI diterima admin, update status ke :", input.Status)
		query := `
			UPDATE performance_invoice
			SET
				status = $1,
				update_at = now(),
				updated_by = 'admin'
			WHERE id = $2;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Status,
			input.ID,
		)

		//  Berhasil Masuk Ke Keuangan Mereka

	} else if input.Status == "Ditolak" {
		log.Println("PI ditolak admin, update status ke :", input.Status)
		log.Println("PI ditolak admin, tambah alasan ke :", input.Reason)

		query := `
			UPDATE performance_invoice
			SET
				status = $1,
				reason = $2,
				update_at = now(),
				updated_by = 'admin'
			WHERE id = $3;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Status,
			input.Reason,
			input.ID,
		)

		//  Berhasil Masuk Ke Keuangan Mereka

	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Update Successfully", "status": true})
}
