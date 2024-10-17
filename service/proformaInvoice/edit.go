package proformaInvoice

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"fmt"
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
	response.ItemDeleted = make([]model.ItemDeleted, len(input.ItemDeleted))

	if len(input.ItemDetailPI) > 0 {

		for i, item := range input.ItemDetailPI {
			log.Println("Data Ke :", i)

			response.ItemDetailPI[i].Kat = item.Kat
			response.ItemDetailPI[i].NamaBarang = item.NamaBarang
			response.ItemDetailPI[i].Quantity = item.Quantity
			response.ItemDetailPI[i].HargaSatuan = item.HargaSatuan
			response.ItemDetailPI[i].Discount = item.Discount
			response.ItemDetailPI[i].Kode = item.Kode
			response.ItemDetailPI[i].Variable = item.Variable
			response.ItemDetailPI[i].Gudang = item.Gudang

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
	response.AlamaCustomer = input.AlamaCustomer
	response.Customer = input.Customer

	now := time.Time{}
	response.Tanggal = utility.FormatTanggal1(now)

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

	input.Status = "Diproses"

	input.TotalRP = utility.RupiahToNumber(input.TotalRP)
	input.SubTotalRP = utility.RupiahToNumber(input.SubTotalRP)
	input.PajakPPNRP = utility.RupiahToNumber(input.PajakPPNRP)

	if input.Divisi == "Radiologi" {
		query := `
			UPDATE performance_invoice
			SET
				customer = $1,
				sub_total = $2,
				status = $3,
				divisi = $4,
				invoice_number = $5,
				due_date = $6,
				pajak = $7,
				total = $8,
				number_si = $9,
				update_at = now(),
				updated_by = 'sales',
				reason = $10,
				rm = $12
			WHERE id = $11;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Customer,
			input.SubTotalRP,
			"Diproses",
			input.Divisi,
			input.InvoiceNumber,
			input.DueDate,
			input.PajakPPNRP,
			input.TotalRP,
			input.NumberSI,
			"",
			input.ID,
			input.RM,
		)

		log.Println("Edit Radiologi")

	} else if input.Divisi == "Ortopedi" {
		query := `
			UPDATE performance_invoice
			SET
				customer = $1,
				sub_total = $2,
				status = $3,
				divisi = $4,
				invoice_number = $5,
				due_date = $6,
				doctor_name = $7,
				patient_name = $8,
				pajak = $9,
				total = $10,
				tanggal_tindakan = $11,
				rm = $12,
				number_si = $13,
				update_at = now(),
				updated_by = 'sales',
				reason = $14
			WHERE id = $15;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Customer,
			input.SubTotalRP,
			input.Status,
			input.Divisi,
			input.InvoiceNumber,
			input.DueDate,
			input.DoctorName,
			input.PatientName,
			input.PajakPPNRP,
			input.TotalRP,
			input.TanggalTindakan,
			input.RM,
			input.NumberSI,
			"",
			input.ID,
		)

		log.Println("Edit Ortopedi")

	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if len(input.ItemDetailPI) > 0 {
		for _, detail := range input.ItemDetailPI {
			if detail.Id == 0 {
				log.Println("Barang Di Tambah ! : ", detail.Id)

				QueryItem := `
					insert into order_items (
						pi_id, "name", quantity, price, discount, 
						sub_total, kat, 
						created_at, created_by, update_at, updated_by, variable, kode, gudang
						) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), 'sales', NOW(), 'sales'), $8, $9, $10
				`

				_, err = tx.Exec(ctx, QueryItem,
					input.ID, detail.NamaBarang, detail.Quantity, detail.HargaSatuan,
					detail.Discount, detail.RPSubTotalItem, detail.Variable, detail.Variable, detail.Kode, detail.Gudang)

				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
					return
				}
			} else if detail.Id > 0 {
				log.Println("Barang Di Edit ! : ", detail.Id)

				queryUPDATE := `
				UPDATE order_items
				SET
					name = $1,
					quantity = $2,
					price = $3,
					discount = $4,
					sub_total = $5,
					kat = $6,
					update_at = $7,
					updated_by = $8,
					variable = $10,
					kode = $11,
					gudang = $12
				WHERE id = $9;
				`

				_, err = tx.Exec(ctx, queryUPDATE,
					detail.NamaBarang,
					detail.Quantity,
					detail.HargaSatuan,
					detail.Discount,
					detail.RPSubTotalItem,
					detail.Variable,
					time.Now(),
					"sales",
					detail.Id,
					detail.Variable,
					detail.Kode,
					detail.Gudang,
				)

				if err != nil {
					err := tx.Rollback(ctx)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
						return
					}
				}
			}
		}
	}

	//  Ketika Berhasil Pengaruhi jumlah Stok

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
		//  Kurangi Jumalh Stok Barang di gudang tujuan
		if len(input.ItemDetailPI) != 0 {
			for _, detail := range input.ItemDetailPI {

				queryGetGudang := `select id from gudang g where nama_gudang = $1 `

				var gudangId int
				err := tx.QueryRow(ctx, queryGetGudang, detail.Gudang).Scan(&gudangId)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query Pengurangan Stock", "status": false})
					return
				}

				log.Println("Gudang Id : ", gudangId)

				queryCekQty := `SELECT qty FROM stock WHERE nama = $1 and gudang_id = $2`
				var currentQty int
				err2 := tx.QueryRow(ctx, queryCekQty, detail.NamaBarang, gudangId).Scan(&currentQty)
				if err2 != nil {
					tx.Rollback(ctx)
					log.Println("Error Qeury : ", err2)

					c.JSON(http.StatusInternalServerError, gin.H{
						"message": fmt.Sprintf("Data di gudang untuk nama barang : <b>" + detail.NamaBarang + " TIDAK ADA!</b>, Silahkan Cek Pada Gudang : <b>" + detail.Gudang + "</b>"),
						"status":  false,
					})
					return
				}

				log.Println("Current Quantity : ", currentQty)

				quantity, _ := strconv.Atoi(detail.Quantity)

				if currentQty-quantity < 0 {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": fmt.Sprintf("Data di gudang untuk nama barang : <b>" + detail.NamaBarang + " TIDAK CUKUP!</b>, Silahkan Cek Pada Gudang : <b>" + detail.Gudang + "</b>"),
						"status":  false,
					})
					return
				}

				queryPenguranganStokc := `
					UPDATE stock
					SET qty = qty - $1
					WHERE nama = $2 and gudang_id = $3
				`

				_, err3 := tx.Exec(ctx, queryPenguranganStokc, quantity, detail.NamaBarang, gudangId)
				if err3 != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query Pengurangan Stock", "status": false})
					return
				}

			}
		}

		log.Println("Proses Duplikasi")

		QueryInsertDuplicate := `
			INSERT INTO performance_invoice_copy
			SELECT *
			FROM performance_invoice
			WHERE id = $1;
		`

		_, err = tx.Exec(context.Background(), QueryInsertDuplicate, input.ID)
		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery3)
			return
		}

		QueryInsertDuplicateItemOrder := `
			INSERT INTO order_items_copy
			SELECT *
			FROM order_items
			WHERE pi_id = $1;
		`

		_, err = tx.Exec(context.Background(), QueryInsertDuplicateItemOrder, input.ID)
		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery3)
			return
		}

		//  Berhasil Masuk Ke Keuangan Mereka

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
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
			return
		}

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
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
			return
		}

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
