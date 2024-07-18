package preOrder

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func ListPO(c *gin.Context) {

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		SELECT 
			COALESCE(id, 0) AS id,
			COALESCE(nama_suplier, '') AS nama_suplier,
			COALESCE(nomor_po, '') AS nomor_po,
			COALESCE(tanggal, '') AS tanggal,
			COALESCE(catatan_po, '') AS catatan_po,
			COALESCE(prepared_by, '') AS prepared_by,
			COALESCE(prepared_jabatan, '') AS prepared_jabatan,
			COALESCE(approved_by, '') AS approved_by,
			COALESCE(approved_jabatan, '') AS approved_jabatan,
			COALESCE(status, '') AS status,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total
		FROM purchase_order order by id asc;
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.PurchaseOrder
	for rows.Next() {
		var res model.PurchaseOrder
		if err := rows.Scan(
			&res.ID,
			&res.NamaSuplier,
			&res.NomorPO,
			&res.Tanggal,
			&res.CatatanPO,
			&res.PreparedBy,
			&res.PreparedJabatan,
			&res.ApprovedBy,
			&res.ApprovedJabatan,
			&res.Status,
			&res.SubTotal,
			&res.Pajak,
			&res.Total,
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
