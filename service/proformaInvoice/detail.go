package proformaInvoice

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
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
			COALESCE(a.customer_id, 0) AS customer_id,
			COALESCE(a.sub_total, '') AS sub_total,
			COALESCE(a.status, '') AS status,
			COALESCE(a.divisi, '') AS divisi,
			COALESCE(a.invoice_number, '') AS invoice_number,
			COALESCE(a.po_number, '') AS po_number,
			COALESCE(a.due_date , '') AS due_date,
			COALESCE(a.doctor_name, '') AS doctor_name,
			COALESCE(a.patient_name, '') AS patient_name,
			COALESCE(a.pajak, '') AS pajak,
			COALESCE(a.total, '') AS total,
			COALESCE(a.tanggal_tindakan, '') AS tanggal_tindakan,
			COALESCE(a.rm, '') AS rm,
			COALESCE(a.number_si, '') AS number_si,
			COALESCE(a.reason, '') AS reason,
			COALESCE(c.nama_company, '') AS nama_company,
			COALESCE(c.address_company , '') AS address_company
		FROM 
			performance_invoice a, customer c where a.customer_id = c.id and a.id = $1
	`

	rows, err := tx.Query(ctx, queryDetailPI, input.ID)
	if err != nil {
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
			&invoice.CustomerID,
			&invoice.SubTotal,
			&invoice.Status,
			&invoice.Divisi,
			&invoice.InvoiceNumber,
			&invoice.PONumber,
			&invoice.DueDate,
			&invoice.DoctorName,
			&invoice.PatientName,
			&invoice.Pajak,
			&invoice.Total,
			&invoice.TanggalTindakan,
			&invoice.RM,
			&invoice.NumberSI,
			&invoice.Reason,
			&invoice.Customer,
			&invoice.AlamaCustomer,
		)

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
			COALESCE(quantity, '') AS quantity,
			COALESCE(price, '') AS price,
			COALESCE(discount, '') AS discount,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(kat, '') AS kat
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
			&ambil.Quantity,
			&ambil.HargaSatuan,
			&ambil.Discount,
			&ambil.SubTotalItem,
			&ambil.Kat,
		)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
			return
		}

		TampungItem = append(TampungItem, ambil)
	}

	invoice.ItemDetailPI = TampungItem

	if input.Export == "YES" {
		utility.ExportPI(c, invoice)
	}

	if invoice.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": invoice, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}
