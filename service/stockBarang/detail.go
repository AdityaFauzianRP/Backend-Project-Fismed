package stockBarang

import (
	"backend_project_fismed/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func Detail(c *gin.Context) {
	// Detail Barang

	var input service.RequestID

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

	query := "select * from stock_items where id = $1"

	rows, err := tx.Query(ctx, query, input.ID)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute tax code query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []service.StockBarang
	for rows.Next() {
		var res service.StockBarang
		if err := rows.Scan(
			&res.ID,
			&res.Name,
			&res.Total,
			&res.Price,
			&res.CreatedAt,
			&res.CreatedBy,
			&res.UpdatedAt,
			&res.UpdatedBy,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Stock Barang", "status": false})
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
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []service.StockBarang{}, "status": true})
	}
}
