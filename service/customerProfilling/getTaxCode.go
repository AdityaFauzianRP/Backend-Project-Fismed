package customerProfilling

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func GetTaxCode(c *gin.Context) {
	type TaxCode struct {
		Id      int    `json:"id"`
		TaxCode string `json:"tax_code"`
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := "select id , tax  from tax_code tc"

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute tax code query", "status": false})
		return
	}
	defer rows.Close()

	var taxCodes []TaxCode
	for rows.Next() {
		var tc TaxCode
		if err := rows.Scan(&tc.Id, &tc.TaxCode); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan tax code", "status": false})
			return
		}
		taxCodes = append(taxCodes, tc)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over tax code rows", "status": false})
		return
	}

	if len(taxCodes) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "tax_codes": taxCodes, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "tax_codes": "No tax codes available", "status": true})
	}

}
