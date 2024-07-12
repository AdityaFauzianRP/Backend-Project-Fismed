package piutang

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

	query := `
		select 
			COALESCE(id, 0) AS id,
			COALESCE(nama, '') AS nama,
			COALESCE(nominal, '') AS nominal,
			COALESCE(amount, '') AS amount,
			tanggal
		from piutang p order by id desc 
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Pemasukan
	var total int = 0
	for rows.Next() {
		var res model.Pemasukan
		if err := rows.Scan(
			&res.Id,
			&res.Nama,
			&res.Nominal,
			&res.Amount,
			&res.Tanggal,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}

		nominal, _ := strconv.Atoi(res.Nominal)
		total = total + nominal
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
