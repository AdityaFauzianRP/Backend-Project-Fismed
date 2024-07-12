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

	queryDetailPI := `
		SELECT 
			COALESCE(id, 0) AS id,
			COALESCE(customer_id, 0) AS customer_id,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(status, '') AS status,
			COALESCE(divisi, '') AS divisi,
			COALESCE(invoice_number, '') AS invoice_number,
			COALESCE(po_number, '') AS po_number,
			COALESCE(due_date , '') AS due_date,
			COALESCE(doctor_name, '') AS doctor_name,
			COALESCE(patient_name, '') AS patient_name,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total,
			COALESCE(tanggal_tindakan, '') AS tanggal_tindakan,
			COALESCE(rm, '') AS rm,
			COALESCE(number_si, '') AS number_si,
			COALESCE(reason, '') AS reason
		FROM 
			performance_invoice
		where id = $1
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
