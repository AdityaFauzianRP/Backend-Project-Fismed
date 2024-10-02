package proformaInvoice

import (
	"backend_project_fismed/constanta"
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

func InquiryPI(c *gin.Context) {
	var input model.ProformaInvoice

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	log.Println("Data Input :", input)

	input.TanggalPI = time.Now().Format("2006-01-02")

	if input.NamaDokter != "" {
		//	Ortopedic
		input.NomorInvoice = utility.GenerateINVNumber(342, "OT")
	} else if input.NamaDokter == "" {
		//	Ortopedic
		input.NomorInvoice = utility.GenerateINVNumber(342, "RG")
	}

	input.NomorSI = "SI/" + utility.GenerateTigaNomor() + "/" + utility.GenerateNomorInvoice()

	var subtotal, total, ppn int

	if len(input.Item) != 0 {
		for i, data := range input.Item {
			harga1, err := strconv.Atoi(data.HargaSatuan)
			if err != nil {
				log.Println("Harga Bukan String !")
				return
			}

			pcs, err := strconv.Atoi(data.Quantity)
			if err != nil {
				log.Println("Quantity Bukan String !")
				return
			}

			input.Item[i].Amount = "Rp. " + utility.FormatRupiah(strconv.Itoa(harga1*pcs-(harga1*pcs*data.Discount/100)))
			input.Item[i].Price = "Rp. " + utility.FormatRupiah(strconv.Itoa(harga1))
			input.Item[i].Price = input.Item[i].HargaSatuan

			subtotal = subtotal + harga1*pcs
		}
	}

	ppn = subtotal * 11 / 100
	total = ppn + subtotal

	input.Total = "Rp. " + utility.FormatRupiah(strconv.Itoa(total))
	input.Pajak = "Rp. " + utility.FormatRupiah(strconv.Itoa(ppn))
	input.Subtotal = "Rp. " + utility.FormatRupiah(strconv.Itoa(subtotal))

	c.JSON(http.StatusOK, gin.H{"message": "Inquiry Performa Invoice Success !", "data": input, "status": true})
}

func PostingPI(c *gin.Context) {
	var input model.ProformaInvoice

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	log.Println("Data Input :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	input.Subtotal = utility.RupiahToNumber(input.Subtotal)
	input.Total = utility.RupiahToNumber(input.Total)
	input.Pajak = utility.RupiahToNumber(input.Pajak)

	var newId int

	queryInsertPI := `
		INSERT INTO performance_invoice (
           customer, sub_total, status, divisi, invoice_number, doctor_name,
           patient_name, created_at, created_by, update_at, updated_by, total, pajak, tanggal_tindakan,
           rm, number_si, alamat_customer
       )
       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id;
	`
	err = tx.QueryRow(ctx, queryInsertPI,
		input.RumahSakit, input.Subtotal, "DIPROSES", input.IDDivisi, input.NomorInvoice, input.NamaDokter, input.NamaPasien,
		time.Now(), "SALES", time.Now(), "SALES", input.Total, input.Pajak, input.TanggalTindakan, input.RM, input.NomorSI, input.Alamat).Scan(&newId)

	if err != nil {
		log.Println("Error Insert PI ! : ", err)
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if len(input.Item) != 0 {
		for _, data := range input.Item {
			discount := strconv.Itoa(data.Discount)

			queryInsertBarangPI := `
				INSERT INTO order_items (
					pi_id, name, quantity, price, discount, sub_total, kode, variable, gudang
				) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			`

			_, err := tx.Exec(ctx, queryInsertBarangPI,
				newId, data.NamaBarang, data.Quantity, data.Price, discount, data.Amount, data.Kode, data.Variable, data.Gudang)

			if err != nil {
				log.Println("Error Insert Barang PI ! : ", err)
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
				return
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		utility.ResponseError(c, constanta.ErrCommit)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data added successfully", "status": true})

}
