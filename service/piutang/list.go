package piutang

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
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
		from pengeluaran p where status = 'PENDING' order by id desc 
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Pengeluaran
	var total int
	for rows.Next() {
		var res model.Pengeluaran
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

func ListPiutang(c *gin.Context) {

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
			COALESCE(pi_id, '') AS pi_id,
			COALESCE(status_pi, '') AS status_pi
		from pemasukan p where status = 'PENDING' order by id desc 
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Pengeluaran
	var total int
	for rows.Next() {
		var res model.Pengeluaran
		if err := rows.Scan(
			&res.Id,
			&res.Nama,
			&res.Nominal,
			&res.Pajak,
			&res.Amount,
			&res.Tanggal,
			&res.Piid,
			&res.Divisi,
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

func Lunas(c *gin.Context) {

	var input model.Pengeluaran

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

	query := `
		UPDATE pengeluaran SET status = $2
		WHERE id = $1
	`

	_, err = tx.Exec(context.Background(), query,
		input.Id,   // $1
		"DITERIMA", // $2
	)
	if err != nil {
		log.Fatalf("Error executing insert: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatalf("Error committing transaction: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Jadi Pengeluaran !", "data": []model.Customer{}, "status": true})

}

func LunasPiutang(c *gin.Context) {

	var input model.Pengeluaran

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

	query := `
		UPDATE pemasukan SET status = $2
		WHERE id = $1
	`

	_, err = tx.Exec(context.Background(), query,
		input.Id,   // $1
		"DITERIMA", // $2
	)
	if err != nil {
		log.Fatalf("Error executing insert: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatalf("Error committing transaction: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Jadi Pengeluaran !", "data": []model.Customer{}, "status": true})

}
