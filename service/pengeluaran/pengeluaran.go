package pengeluaran

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		select 
			COALESCE(id, 0) AS id,
			COALESCE(nama, '') AS nama,
			COALESCE(sub_total , '') AS sub_total,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total,
			tanggal,
			COALESCE(po_id, '') AS po_id
		from pengeluaran p where status = 'DITERIMA' order by id desc 
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Pemasukan
	var total int
	for rows.Next() {
		var res model.Pemasukan
		if err := rows.Scan(
			&res.Id,
			&res.Nama,
			&res.Nominal,
			&res.Pajak,
			&res.Amount,
			&res.Tanggal,
			&res.Piid,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Amount, _ := strconv.Atoi(res.Amount)
		total = total + Amount

		res.Nominal = "Rp. " + utility.FormatRupiah(res.Nominal)
		res.Pajak = "Rp. " + utility.FormatRupiah(res.Pajak)
		res.Amount = "Rp. " + utility.FormatRupiah(res.Amount)

		Responses = append(Responses, res)
	}

	total_pemasukan := utility.FormatRupiah(strconv.Itoa(total))

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data Ditemukan !",
			"data":    Responses,
			"total":   "Rp. " + total_pemasukan,
			"status":  true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}
