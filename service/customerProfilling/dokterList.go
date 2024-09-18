package customerProfilling

import (
	"backend_project_fismed/model"
	"backend_project_fismed/service/proformaInvoice"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func DokterList(c *gin.Context) {
	ctx := context.Background()
	tx, err := proformaInvoice.DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("Get Rumash Sakit !")
	query := `
		SELECT DISTINCT ON (nama_dokter) 
			COALESCE(nama_dokter, 'default_name') AS name
		FROM customer
		ORDER BY nama_dokter ASC;
		`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.DocktorName
	for rows.Next() {
		var res model.DocktorName
		if err := rows.Scan(
			&res.Nama,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Responses = append(Responses, res)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Responses, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}
