package price_list

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func PriceList(c *gin.Context) {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	queryList := `
		select 
			nama_rumah_sakit ,
			kode ,
			variable ,
			nama ,
			price ,
			diskon
		from price_list pl 
	`

	row, err := tx.Query(ctx, queryList)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}

	defer row.Close()

	var Data []model.Price

	for row.Next() {
		var ambil model.Price
		err := row.Scan(
			&ambil.NamarumahSakit,
			&ambil.Kode,
			&ambil.Variable,
			&ambil.Nama,
			&ambil.Price,
			&ambil.Diskon,
		)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Stock Barang  123", "status": false})
			return
		}
		Data = append(Data, ambil)
	}

	if err := row.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Data) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Data, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}
}

func ListByCustomer(c *gin.Context) {

	var input model.RequestID

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

	queryList := `
		select 
			nama_rumah_sakit ,
			kode ,
			variable ,
			nama ,
			price ,
			diskon
		from price_list pl where nama_rumah_sakit = $1
	`

	row, err := tx.Query(ctx, queryList, input.Nama)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}

	defer row.Close()

	var Data []model.Price

	for row.Next() {
		var ambil model.Price
		err := row.Scan(
			&ambil.NamarumahSakit,
			&ambil.Kode,
			&ambil.Variable,
			&ambil.Nama,
			&ambil.Price,
			&ambil.Diskon,
		)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Stock Barang  123", "status": false})
			return
		}
		Data = append(Data, ambil)
	}

	if err := row.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Data) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Data, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}
}

func SetPrice(c *gin.Context) {
	var input model.Set

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

	queryInsert := `
		INSERT INTO price_list (nama_rumah_sakit, kode, variable, nama, diskon, price) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	log.Println("Set To DB Start")

	if len(input.Input) != 0 {
		for _, set := range input.Input {
			_, err := tx.Exec(context.Background(), queryInsert, set.NamarumahSakit, set.Kode, set.Variable, set.Nama, set.Diskon, set.Price)
			if err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
				return
			}
		}
	}

	log.Println("Set To DB Done")

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Set successfully", "status": true})
}
