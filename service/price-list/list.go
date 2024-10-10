package price_list

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
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

	cek := 0

	queryCek := `
		select count(nama_rumah_sakit)from price_list pl where nama_rumah_sakit = $1;
	`
	err = tx.QueryRow(ctx, queryCek, input.Nama).Scan(&cek)

	if cek == 0 {
		log.Println("Customer Baru, Hargannya blm di set !")
		queryList := `
			SELECT 
				a.variable,
				a.nama,
				a.kode,
				CAST(a.price AS VARCHAR) AS price
			FROM stock a
			GROUP BY 
				a.variable,
				a.nama,
				a.kode,
				a.price
			ORDER BY a.nama;
		`

		row, err := tx.Query(ctx, queryList)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
			return
		}

		defer row.Close()

		var Data []model.Price
		hitung := 0

		for row.Next() {
			var ambil model.Price
			err := row.Scan(
				&ambil.Variable,
				&ambil.Nama,
				&ambil.Kode,
				&ambil.Price,
			)

			if err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Data", "status": false})
				return
			}
			ambil.Name = ambil.Nama
			ambil.NamarumahSakit = input.Nama
			ambil.Diskon = 0
			ambil.Added = "0"

			ambil.ID = hitung + 1
			hitung = hitung + 1
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

	} else {
		log.Println("Customer Lama, Sudah Set Harga !")

		queryList := `
		select 
			nama_rumah_sakit ,
			kode ,
			variable ,
			nama ,
			price ,
			diskon ,
			added
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
		hitung := 0

		for row.Next() {
			var ambil model.Price
			err := row.Scan(
				&ambil.NamarumahSakit,
				&ambil.Kode,
				&ambil.Variable,
				&ambil.Nama,
				&ambil.Price,
				&ambil.Diskon,
				&ambil.Added,
			)
			ambil.ID = hitung + 1
			hitung = hitung + 1
			ambil.Name = ambil.Nama

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

}

func SetPrice(c *gin.Context) {
	var input model.Set

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
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
		INSERT INTO price_list (nama_rumah_sakit, kode, variable, nama, diskon, price, added) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	queryUpdate := `
		UPDATE price_list
		SET 
			nama_rumah_sakit = $1,
			kode = $2,
			variable = $3,
			nama = $4,
			diskon = $5,
			price = $6,
			added = $7
		WHERE nama_rumah_sakit = $8 AND nama = $9
	`

	log.Println("Set To DB Start")

	if len(input.Input) != 0 {

		for _, set := range input.Input {

			if set.Added == "0" {
				log.Println("Tambah Ke Price List")
				var id int
				err := tx.QueryRow(ctx, "SELECT MIN(id) FROM customer WHERE nama_perusahaan = $1;", set.NamarumahSakit).Scan(&id)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
					return
				}

				set.Kode = utility.GenerateKode(id, set.NamarumahSakit)

				_, err = tx.Exec(ctx, queryInsert, set.NamarumahSakit, set.Kode, set.Variable, set.Nama, set.Diskon, set.Price, "1")
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
					return
				}
			} else {
				log.Println("Update Ke Price List")
				_, err := tx.Exec(ctx, queryUpdate, set.NamarumahSakit, set.Kode, set.Variable, set.Nama, set.Diskon, set.Price, "1", set.NamarumahSakit, set.Nama)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
					return
				}
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
