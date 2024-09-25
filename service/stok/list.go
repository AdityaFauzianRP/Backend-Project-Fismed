package stok

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func ListByGudangId(c *gin.Context) {

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
			a.variable,
			a.nama,
			a.qty,
			a.price,
			a.kode, 
			g.nama_gudang
		from stock a , gudang g where gudang_id = $1 and a.gudang_id  = g.id 
	`

	row, err := tx.Query(ctx, queryList, input.ID)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}

	defer row.Close()

	var Data []model.Stock

	for row.Next() {
		var ambil model.Stock
		err := row.Scan(
			&ambil.Variable,
			&ambil.Nama,
			&ambil.Qty,
			&ambil.Harga,
			&ambil.Kode,
			&ambil.NamaGudang,
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

func ListBarang(c *gin.Context) {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	queryList := `
		select 
			a.variable,
			a.nama,
			a.qty,
			a.price,
			a.kode,
			g.nama_gudang
		from stock a, gudang g where a.gudang_id = g.id
	`

	row, err := tx.Query(ctx, queryList)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}

	defer row.Close()

	var Data []model.Stock

	for row.Next() {
		var ambil model.Stock
		err := row.Scan(
			&ambil.Variable,
			&ambil.Nama,
			&ambil.Qty,
			&ambil.Harga,
			&ambil.Kode,
			&ambil.NamaGudang,
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

func ListBarangProses(c *gin.Context) {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	queryList := `
		SELECT 
			a.variable,
			a.nama,
			SUM(a.qty) AS total_qty,
			a.kode,
			a.price
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

	var Data []model.Stock

	for row.Next() {
		var ambil model.Stock
		err := row.Scan(
			&ambil.Variable,
			&ambil.Nama,
			&ambil.Qty,
			&ambil.Kode,
			&ambil.Harga,
		)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Data", "status": false})
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
