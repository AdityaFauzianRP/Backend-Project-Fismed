package gudang

import (
	"backend_project_fismed/model"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GudangListData(c *gin.Context) {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	queryList := `
		SELECT id, nama_gudang , lokasi  FROM gudang order by id 
	`

	row, err := tx.Query(ctx, queryList)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}

	defer row.Close()

	var Data []model.Gudang

	for row.Next() {
		var ambil model.Gudang
		err := row.Scan(
			&ambil.Id,
			&ambil.Nama,
			&ambil.Lokasi,
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

func TambahGudang(c *gin.Context) {

	var input model.Gudang

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

	queryTambahGudang := `
		INSERT INTO gudang (
			nama_gudang,
			lokasi
		) VALUES (
			$1, $2
		)
	`

	_, err = tx.Exec(ctx, queryTambahGudang,
		input.Nama,
		input.Lokasi,
	)

	if err != nil {
		tx.Rollback(ctx)
		log.Println("Insert failed:", err)
		c.JSON(400, gin.H{"message": err})
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(400, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditambahkan !", "status": true})

}

func HapusGudang(c *gin.Context) {
	// Ambil ID dari parameter URL
	var input model.Gudang

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

	queryHapusGudang := `
		DELETE FROM gudang
		WHERE id = $1
	`

	_, err = tx.Exec(ctx, queryHapusGudang, input.Id)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("Delete failed:", err)
		c.JSON(400, gin.H{"message": err})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(400, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus!", "status": true})
}

func AkunListData(c *gin.Context) {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	queryList := `
		SELECT u.id, u.username , uc."role"  FROM users u , user_category uc where u.role_id = uc.id  order by id  
	`

	row, err := tx.Query(ctx, queryList)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}

	defer row.Close()

	var Data []model.Akun

	for row.Next() {
		var ambil model.Akun
		err := row.Scan(
			&ambil.Id,
			&ambil.Nama,
			&ambil.Password,
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

func TambahAkun(c *gin.Context) {

	var input model.Akun

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

	hasher := md5.New()
	hasher.Write([]byte(input.Password))
	hashedInputPassword := hex.EncodeToString(hasher.Sum(nil))

	roleint, _ := strconv.Atoi(input.RoleID)

	queryTambahAkun := `
		INSERT INTO users (
			username,
			password,
			role_id,
			created_at ,
			created_by ,
			updated_at ,
			updated_by ,
		    token
		) VALUES (
			$1, $2 ,$3, $4, $5, $6, $7, $8
		)
	`

	_, err = tx.Exec(ctx, queryTambahAkun,
		input.Nama,
		hashedInputPassword,
		roleint,
		time.Now(),
		"SUPER ADMIN",
		time.Now(),
		"SUPER ADMIN",
		"Empty",
	)

	if err != nil {
		tx.Rollback(ctx)
		log.Println("Insert failed:", err)
		c.JSON(400, gin.H{"message": err})
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(400, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditambahkan !", "status": true})

}

func HapusAkun(c *gin.Context) {
	// Ambil ID dari parameter URL
	var input model.Gudang

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

	queryHapusGudang := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err = tx.Exec(ctx, queryHapusGudang, input.Id)
	if err != nil {
		tx.Rollback(ctx)
		log.Println("Delete failed:", err)
		c.JSON(400, gin.H{"message": err})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		c.JSON(400, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus!", "status": true})
}
