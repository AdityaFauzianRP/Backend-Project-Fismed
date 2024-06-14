package stockBarang

import (
	"backend_project_fismed/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func Edit(c *gin.Context) {
	// Edit Stock Barang
	var input service.StockBarang

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

	query := `
		UPDATE stock_items
		SET name = $1,
			total = $2,
			price = $3
		WHERE id = $4;
	`

	_, err = tx.Exec(ctx, query,
		input.Name, input.Total, input.Price, input.ID)

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully", "status": true})
}
