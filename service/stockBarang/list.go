package stockBarang

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func List(c *gin.Context) {
	//	List Semua Barang dan jumlah Stocknnya

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := "SELECT id, name, total, price, created_at, created_by, updated_at, updated_by FROM public.stock_items;"

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.StockBarang
	for rows.Next() {
		var res model.StockBarang
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
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}
}
