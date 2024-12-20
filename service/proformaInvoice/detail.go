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
	"strings"
)

func DetailPI(c *gin.Context) {
	// Detail Data PI

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

	queryDetailPI := `
		SELECT 
			COALESCE(a.id, 0) AS id,
			COALESCE(a.customer, '') AS customer,
			COALESCE(a.sub_total, '') AS sub_total,
			COALESCE(a.status, '') AS status,
			COALESCE(a.divisi, '') AS divisi,
			COALESCE(a.invoice_number, '') AS invoice_number,
			COALESCE(a.due_date , '') AS due_date,
			COALESCE(a.doctor_name, '') AS doctor_name,
			COALESCE(a.patient_name, '') AS patient_name,
			COALESCE(a.pajak, '') AS pajak,
			COALESCE(a.total, '') AS total,
			COALESCE(a.tanggal_tindakan, '') AS tanggal_tindakan,
			COALESCE(a.rm, '') AS rm,
			COALESCE(a.number_si, '') AS number_si,
			COALESCE(a.reason, '') AS reason,
			COALESCE(a.alamat_customer , '') AS alamat_customer
		FROM 
			performance_invoice a where a.id = $1
	`

	rows, err := tx.Query(ctx, queryDetailPI, input.ID)
	if err != nil {
		log.Println("Error Detail PI ! : ", err)
		return
	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var invoice model.PerformanceInvoiceDetail

	for rows.Next() {
		err := rows.Scan(
			&invoice.ID,
			&invoice.Customer,
			&invoice.SubTotal,
			&invoice.Status,
			&invoice.Divisi,
			&invoice.InvoiceNumber,
			&invoice.DueDate,
			&invoice.DoctorName,
			&invoice.PatientName,
			&invoice.Pajak,
			&invoice.Total,
			&invoice.TanggalTindakan,
			&invoice.RM,
			&invoice.NumberSI,
			&invoice.Reason,
			&invoice.AlamaCustomer,
		)

		invoice.Pajak = "Rp. " + utility.FormatRupiah(invoice.Pajak)
		invoice.Total = "Rp. " + utility.FormatRupiah(invoice.Total)
		invoice.SubTotal = "Rp. " + utility.FormatRupiah(invoice.SubTotal)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}
	}

	QueryItem := `
		select 
			id , 
			COALESCE("name", '') AS "name",
			COALESCE(variable , '') AS variable,
			COALESCE(quantity, '') AS quantity,
			COALESCE(price, '') AS price,
			COALESCE(discount, '') AS discount,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(kode, '') AS kode,
			COALESCE(gudang, '') AS gudang
		from order_items oi where pi_id = $1
	`

	rows, err = tx.Query(ctx, QueryItem, invoice.ID)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var TampungItem []model.ResItemDetail

	for rows.Next() {
		var ambil model.ResItemDetail

		err := rows.Scan(
			&ambil.Id,
			&ambil.NamaBarang,
			&ambil.Variable,
			&ambil.Quantity,
			&ambil.HargaSatuan,
			&ambil.Discount,
			&ambil.SubTotalItem,
			&ambil.Kat,
			&ambil.Gudang,
		)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}
		ambil.Kode = ambil.Kat

		TampungItem = append(TampungItem, ambil)
	}

	if len(TampungItem) != 0 {
		invoice.ItemDetailPI = TampungItem

	} else {
		invoice.ItemDetailPI = []model.ResItemDetail{}
	}

	if input.Export == "YES" {
		utility.ExportPI(c, invoice)
	}

	if invoice.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": invoice, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}

func DetailPISOLogistik(c *gin.Context) {
	// Detail Data PI

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

	queryDetailPI := `
		SELECT 
			COALESCE(a.id, 0) AS id,
			COALESCE(a.customer, '') AS customer,
			COALESCE(a.sub_total, '') AS sub_total,
			COALESCE(a.status, '') AS status,
			COALESCE(a.divisi, '') AS divisi,
			COALESCE(a.invoice_number, '') AS invoice_number,
			COALESCE(a.due_date , '') AS due_date,
			COALESCE(a.doctor_name, '') AS doctor_name,
			COALESCE(a.patient_name, '') AS patient_name,
			COALESCE(a.pajak, '') AS pajak,
			COALESCE(a.total, '') AS total,
			COALESCE(a.tanggal_tindakan, '') AS tanggal_tindakan,
			COALESCE(a.rm, '') AS rm,
			COALESCE(a.number_si, '') AS number_si,
			COALESCE(a.reason, '') AS reason,
			COALESCE(a.alamat_customer , '') AS alamat_customer
		FROM 
			performance_invoice a where a.id = $1
	`

	rows, err := tx.Query(ctx, queryDetailPI, input.ID)
	if err != nil {
		log.Println("Error Detail PI ! : ", err)
		return
	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var invoice model.PerformanceInvoiceDetail

	for rows.Next() {
		err := rows.Scan(
			&invoice.ID,
			&invoice.Customer,
			&invoice.SubTotal,
			&invoice.Status,
			&invoice.Divisi,
			&invoice.InvoiceNumber,
			&invoice.DueDate,
			&invoice.DoctorName,
			&invoice.PatientName,
			&invoice.Pajak,
			&invoice.Total,
			&invoice.TanggalTindakan,
			&invoice.RM,
			&invoice.NumberSI,
			&invoice.Reason,
			&invoice.AlamaCustomer,
		)

		invoice.Pajak = "Rp. " + utility.FormatRupiah(invoice.Pajak)
		invoice.Total = "Rp. " + utility.FormatRupiah(invoice.Total)
		invoice.SubTotal = "Rp. " + utility.FormatRupiah(invoice.SubTotal)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}
	}

	QueryItem := `
		select 
			id , 
			COALESCE("name", '') AS "name",
			COALESCE(variable , '') AS variable,
			COALESCE(quantity, '') AS quantity,
			COALESCE(price, '') AS price,
			COALESCE(discount, '') AS discount,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(kode, '') AS kode,
			COALESCE(gudang, '') AS gudang
		from order_items oi where pi_id = $1
	`

	rows, err = tx.Query(ctx, QueryItem, invoice.ID)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var TampungItem []model.ResItemDetail

	for rows.Next() {
		var ambil model.ResItemDetail

		err := rows.Scan(
			&ambil.Id,
			&ambil.NamaBarang,
			&ambil.Variable,
			&ambil.Quantity,
			&ambil.HargaSatuan,
			&ambil.Discount,
			&ambil.SubTotalItem,
			&ambil.Kat,
			&ambil.Gudang,
		)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}
		ambil.Kode = ambil.Kat

		TampungItem = append(TampungItem, ambil)
	}

	if len(TampungItem) != 0 {
		invoice.ItemDetailPI = TampungItem

	} else {
		invoice.ItemDetailPI = []model.ResItemDetail{}
	}

	if input.Export == "YES" {
		utility.ExportPI(c, invoice)
	}

	if invoice.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": invoice, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}

func DetailPISO(c *gin.Context) {
	// Detail Data PI

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

	queryDetailPI := `
		SELECT 
			COALESCE(a.id, 0) AS id,
			COALESCE(a.customer, '') AS customer,
			COALESCE(a.sub_total, '') AS sub_total,
			COALESCE(a.status, '') AS status,
			COALESCE(a.divisi, '') AS divisi,
			COALESCE(a.invoice_number, '') AS invoice_number,
			COALESCE(a.due_date , '') AS due_date,
			COALESCE(a.doctor_name, '') AS doctor_name,
			COALESCE(a.patient_name, '') AS patient_name,
			COALESCE(a.pajak, '') AS pajak,
			COALESCE(a.total, '') AS total,
			COALESCE(a.tanggal_tindakan, '') AS tanggal_tindakan,
			COALESCE(a.rm, '') AS rm,
			COALESCE(a.number_si, '') AS number_si,
			COALESCE(a.reason, '') AS reason,
			COALESCE(a.alamat_customer , '') AS alamat_customer,
			TO_CHAR(update_at , 'DD-MM-YYYY') AS update_at_string
		FROM 
			performance_invoice_copy a where a.id = $1
	`

	rows, err := tx.Query(ctx, queryDetailPI, input.ID)
	if err != nil {
		log.Println("Error Detail PI ! : ", err)
		return
	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var invoice model.PerformanceInvoiceDetail

	for rows.Next() {
		err := rows.Scan(
			&invoice.ID,
			&invoice.Customer,
			&invoice.SubTotal,
			&invoice.Status,
			&invoice.Divisi,
			&invoice.InvoiceNumber,
			&invoice.DueDate,
			&invoice.DoctorName,
			&invoice.PatientName,
			&invoice.Pajak,
			&invoice.Total,
			&invoice.TanggalTindakan,
			&invoice.RM,
			&invoice.NumberSI,
			&invoice.Reason,
			&invoice.AlamaCustomer,
			&invoice.Tanggal,
		)

		invoice.Pajak = "Rp. " + utility.FormatRupiah(invoice.Pajak)
		invoice.Total = "Rp. " + utility.FormatRupiah(invoice.Total)
		invoice.SubTotal = "Rp. " + utility.FormatRupiah(invoice.SubTotal)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}
	}

	QueryItem := `
		select 
			id , 
			COALESCE("name", '') AS "name",
			COALESCE(variable , '') AS variable,
			COALESCE(quantity, '') AS quantity,
			COALESCE(price, '') AS price,
			COALESCE(discount, '') AS discount,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(kode, '') AS kode,
			COALESCE(gudang, '') AS gudang
		from order_items_copy oi where pi_id = $1
	`

	rows, err = tx.Query(ctx, QueryItem, invoice.ID)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var TampungItem []model.ResItemDetail

	for rows.Next() {
		var ambil model.ResItemDetail

		err := rows.Scan(
			&ambil.Id,
			&ambil.NamaBarang,
			&ambil.Variable,
			&ambil.Quantity,
			&ambil.HargaSatuan,
			&ambil.Discount,
			&ambil.SubTotalItem,
			&ambil.Kat,
			&ambil.Gudang,
		)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}
		ambil.Kode = ambil.Kat

		TampungItem = append(TampungItem, ambil)
	}

	if len(TampungItem) != 0 {
		invoice.ItemDetailPI = TampungItem

	} else {
		invoice.ItemDetailPI = []model.ResItemDetail{}
	}

	totalstr := utility.RupiahToNumber(invoice.Total)
	total, err := strconv.Atoi(totalstr)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert total to integer", "status": false})
		return
	}

	invoice.Terbilang = terbilang(total)

	if invoice.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": invoice, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}

var angkaKeKata = map[int]string{
	0: "", 1: "Satu", 2: "Dua", 3: "Tiga", 4: "Empat",
	5: "Lima", 6: "Enam", 7: "Tujuh", 8: "Delapan", 9: "Sembilan",
}

var grupAngka = []string{"", "Ribu", "Juta", "Miliar", "Triliun"}

// Fungsi utama konversi angka ke teks
func terbilang(angka int) string {
	if angka == 0 {
		return "Nol Rupiah"
	}

	hasil := ""
	grup := 0

	for angka > 0 {
		ratusan := angka % 1000
		if ratusan > 0 {
			hasil = grupTeks(ratusan) + " " + grupAngka[grup] + " " + hasil
		}
		angka /= 1000
		grup++
	}

	return strings.TrimSpace(hasil) + " Rupiah"
}

// Konversi ratusan ke teks
func grupTeks(angka int) string {
	ratus := angka / 100
	puluhan := angka % 100
	hasil := ""

	if ratus > 0 {
		if ratus == 1 {
			hasil += "Seratus "
		} else {
			hasil += angkaKeKata[ratus] + " Ratus "
		}
	}

	if puluhan > 0 {
		if puluhan < 10 {
			hasil += angkaKeKata[puluhan]
		} else if puluhan < 20 {
			if puluhan == 11 {
				hasil += "Sebelas"
			} else {
				hasil += angkaKeKata[puluhan%10] + " Belas"
			}
		} else {
			hasil += angkaKeKata[puluhan/10] + " Puluh " + angkaKeKata[puluhan%10]
		}
	}

	return strings.TrimSpace(hasil)
}
