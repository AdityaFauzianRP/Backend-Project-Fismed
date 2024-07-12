package preOrder

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func Edit_Admin(c *gin.Context) {
	var input model.PurchaseOrder
	//var response model.PurchaseOrder

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	}

	log.Println("Data Input :", input)

}

func Edit_Finance(c *gin.Context) {
	var input model.PurchaseOrder
	//var response model.PurchaseOrder

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	}

	log.Println("Data Input :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	if input.Status == "DITOLAK" {
		query := `
			UPDATE purchase_order
			SET
				status = $1,
				reason = $2,
				updated_at = now(),
				updated_by = 'admin'
			WHERE id = $3;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Status,
			input.Reason,
			input.ID,
		)

		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery1)
			return
		}

		if err := tx.Commit(ctx); err != nil {
			utility.ResponseError(c, constanta.ErrCommit)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data Edit Success !", "status": true})
	}

	if input.Status == "DITERIMA" {
		query := `
			UPDATE purchase_order
			SET
				status = $1,
				updated_at = now(),
				updated_by = 'admin'
			WHERE id = $2;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Status,
			input.ID,
		)

		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery2)
			return
		}

		// Masukan Ke Table Pengeluaran sebagai catatan

		QueryPengeluaran := `
				INSERT INTO pengeluaran (
					nama, 
					sub_total,
					pajak,
					total,
					tanggal  
				) VALUES ($1, $2, $3, $4, $5)`

		_, err = tx.Exec(context.Background(), QueryPengeluaran, input.NamaSuplier, input.SubTotal, input.Pajak, input.Total, input.Tanggal)
		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery3)
			return
		}

		if err := tx.Commit(ctx); err != nil {
			utility.ResponseError(c, constanta.ErrCommit)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data Edit Success !", "status": true})
	}
}
