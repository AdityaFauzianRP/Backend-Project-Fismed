package proformaInvoice

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func GetAllList(c *gin.Context) {
	//	Get All List PI
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		select 
			COALESCE(a.id, 0) AS id, 
			COALESCE(a.customer_id, 0) AS customer_id, 
			COALESCE(a.status, '') AS status, 
			COALESCE(a.divisi, '') AS divisi, 
			COALESCE(a.invoice_number, '') AS invoice_number, 
			COALESCE(a.po_number, '') AS po_number,  
			COALESCE(a.sub_total, '') AS sub_total, 
			COALESCE(a.pajak, '') AS pajak, 
			COALESCE(a.total, '') AS total, TO_CHAR(
			COALESCE(a.created_at, '1970-01-01 00:00:00'::timestamp), 'YYYY-MM-DD') AS created_at, 
			COALESCE(a.created_by, '') AS created_by, TO_CHAR(
			COALESCE(a.update_at, '1970-01-01 00:00:00'::timestamp), 'YYYY-MM-DD') AS update_at, 
			COALESCE(a.updated_by, '') AS updated_by ,
			C.nama_company 
		from performance_invoice a, customer c where A.customer_id = C.id  ORDER BY id;

	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.PerformanceInvoice
	for rows.Next() {
		var res model.PerformanceInvoice
		if err := rows.Scan(
			&res.ID,
			&res.CustomerID,
			&res.Status,
			&res.Divisi,
			&res.InvoiceNumber,
			&res.PONumber,
			&res.SubTotal,
			&res.Pajak,
			&res.Total,
			&res.CreatedAt,
			&res.CreatedBy,
			&res.UpdateAt,
			&res.UpdatedBy,
			&res.NamaCompany,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Responses = append(Responses, res)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Responses, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.PerformanceInvoice{}, "status": true})
	}
}
