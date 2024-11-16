package stok

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
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
	var id int = 1

	for row.Next() {
		var ambil model.Stock
		ambil.Id = id
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
		ambil.Name = ambil.Nama
		ambil.Price = ambil.Harga
		Data = append(Data, ambil)
		id = id + 1
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

func ListBarangTerjual(c *gin.Context) {
	var input model.RequestID
	var startDate, endDate string

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

	querySelect := "SELECT a.name, SUM(CAST(a.quantity AS INTEGER)) AS total_quantity, TO_CHAR(b.created_at, 'DD-MM-YYYY'), b.customer FROM order_items a JOIN performance_invoice b ON a.pi_id = b.id"

	if input.StartDate != "" {
		startDate = " WHERE b.created_at >= '" + input.StartDate + "' "
		querySelect = querySelect + startDate
	}

	if input.EndDate != "" {
		endDatePlusOne := input.EndDate[:10] + " 23:59:59"

		endDate = "and b.created_at < '" + endDatePlusOne + "' "
		querySelect = querySelect + endDate
	}

	if input.NamaBarang != "" {
		querySelect = querySelect + " AND a.name = '" + input.NamaBarang + "' "
	}

	if input.Nama != "" {
		querySelect = querySelect + " AND b.customer = $1 "
	}

	querySelect = querySelect + " GROUP BY a.name, TO_CHAR(b.created_at, 'DD-MM-YYYY'), b.customer ORDER BY a.name;"

	log.Println("Query Select : ", querySelect)

	var Data []model.BarangTerjual
	var totalBarangTerjual int

	if input.Nama != "" {
		row, err := tx.Query(ctx, querySelect, input.Nama)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
			return
		}

		defer row.Close()

		for row.Next() {
			var ambil model.BarangTerjual
			var quantity int
			err := row.Scan(
				&ambil.NamaBarang,
				&quantity,
				&ambil.Tanggal,
				&ambil.Customer,
			)

			totalBarangTerjual += quantity
			ambil.Qty = strconv.Itoa(quantity)

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
	} else {
		row, err := tx.Query(ctx, querySelect)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
			return
		}

		defer row.Close()

		for row.Next() {
			var ambil model.BarangTerjual
			var quantity int
			err := row.Scan(
				&ambil.NamaBarang,
				&quantity,
				&ambil.Tanggal,
				&ambil.Customer,
			)

			totalBarangTerjual += quantity
			ambil.Qty = strconv.Itoa(quantity)

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
	}

	if len(Data) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Data, "status": true, "total_barang_terjual": totalBarangTerjual})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}
