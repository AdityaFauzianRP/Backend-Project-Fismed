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

func RumahSakitList(c *gin.Context) {
	ctx := context.Background()
	tx, err := proformaInvoice.DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("Get Rumash Sakit !")
	query := `
		SELECT DISTINCT ON (nama_perusahaan) 
			id, 
			COALESCE(nama_perusahaan, 'default_name') AS name, 
			COALESCE(address_perusahaan, 'default_address') AS address_company
		FROM customer
		ORDER BY nama_perusahaan, id ASC;
		`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Customer
	for rows.Next() {
		var res model.Customer
		if err := rows.Scan(
			&res.ID,
			&res.Name,
			&res.AddressCompany,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Responses = append(Responses, res)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Responses, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}

func SupplierList(c *gin.Context) {

	ctx := context.Background()
	tx, err := proformaInvoice.DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("Get Rumash Sakit !")
	query := `
		SELECT DISTINCT ON (nama_perusahaan) 
			id, 
			COALESCE(nama_perusahaan, 'default_name') AS name, 
			COALESCE(address_perusahaan, 'default_address') AS address_company
		FROM customer where kategori_divisi = '3'
		ORDER BY nama_perusahaan, id ASC;
		`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Customer
	for rows.Next() {
		var res model.Customer
		if err := rows.Scan(
			&res.ID,
			&res.Name,
			&res.AddressCompany,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Responses = append(Responses, res)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Responses, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}

func RumahSakitListS(c *gin.Context) {
	ctx := context.Background()
	tx, err := proformaInvoice.DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("Get Rumash Sakit !")
	query := `
		SELECT DISTINCT ON (nama_perusahaan) 
			id, 
			COALESCE(nama_perusahaan, 'default_name') AS name, 
			COALESCE(address_perusahaan, 'default_address') AS address_company
		FROM customer where kategori_divisi = '3'
		ORDER BY nama_perusahaan, id ASC;
		`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Customer
	for rows.Next() {
		var res model.Customer
		if err := rows.Scan(
			&res.ID,
			&res.Name,
			&res.AddressCompany,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Responses = append(Responses, res)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Responses, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}

func RumahSakitListC(c *gin.Context) {
	ctx := context.Background()
	tx, err := proformaInvoice.DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("Get Rumash Sakit !")
	query := `
		SELECT DISTINCT ON (nama_perusahaan) 
			id, 
			COALESCE(nama_perusahaan, 'default_name') AS name, 
			COALESCE(address_perusahaan, 'default_address') AS address_company
		FROM customer where kategori_divisi != '3'
		ORDER BY nama_perusahaan, id ASC;
		`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Customer
	for rows.Next() {
		var res model.Customer
		if err := rows.Scan(
			&res.ID,
			&res.Name,
			&res.AddressCompany,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Responses = append(Responses, res)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over Stock Barang rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Responses, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}
