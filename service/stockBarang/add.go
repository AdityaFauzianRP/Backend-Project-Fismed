package stockBarang

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func Add(c *gin.Context) {
	// Tambah Jenis Barang

	var input model.StockBarang

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

	query := `INSERT INTO stock_items (name, total, price, created_at, created_by, updated_at, updated_by)
	VALUES (
		$1, $2, $3, NOW(), 'admin', NOW(), 'admin'
	);`

	_, err = tx.Exec(ctx, query,
		input.Name, input.Total, input.Price)

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data added successfully", "status": true})
}
