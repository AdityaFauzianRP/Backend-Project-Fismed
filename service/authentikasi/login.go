package authentikasi

import (
	"backend_project_fismed/utility"
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Login(c *gin.Context) {
	type Pengguna struct {
		ID           int    `json:"id"`
		Username     string `json:"username"`
		RolePengguna string `json:"role_name"`
		RoleID       int    `json:"role_id"`
		Token        string `json:"token"`
	}

	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input Request

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

	if input.Username == "" {
		c.JSON(400, gin.H{
			"status":  false,
			"message": "Username Kosong",
		})
		err = tx.Commit(ctx)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	if input.Password == "" {
		c.JSON(400, gin.H{
			"status":  false,
			"message": "Password Kosong",
		})
		err = tx.Commit(ctx)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	var validasi_pass string
	var idp int

	cek_password := `
		select password, id from users u where username = $1
	`
	err = tx.QueryRow(ctx, cek_password, input.Username).Scan(&validasi_pass, &idp)
	if err != nil {
		log.Println("Data User Not Found")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Data User Not Found",
		})
		return
	}

	hasher := md5.New()
	hasher.Write([]byte(input.Password))
	hashedInputPassword := hex.EncodeToString(hasher.Sum(nil))

	if hashedInputPassword != validasi_pass {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Password Salah, Login Gagal",
		})
		err = tx.Commit(ctx)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	log.Println("Login Successfully, Send Data Users For Frontend")

	query_pengguna := `
	select a.id, a.username , a."token" , a.role_id , b."role" from users a, user_category b where a.id = $1 and a.role_id = b.id 
	`
	row1, err := tx.Query(ctx, query_pengguna, idp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		err = tx.Commit(ctx)
		if err != nil {
			panic(err.Error())
		}
		return
	}

	defer row1.Close()

	var Tampung_pengguna []Pengguna

	for row1.Next() {
		var ambil Pengguna
		err := row1.Scan(
			&ambil.ID,
			&ambil.Username,
			&ambil.Token,
			&ambil.RoleID,
			&ambil.RolePengguna,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		Tampung_pengguna = append(Tampung_pengguna, ambil)
	}

	if len(Tampung_pengguna) > 0 {
		Tampung_pengguna[0].Token, err = utility.GenerateToken(Tampung_pengguna[0].ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed Generated Token"})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		err = utility.UpdateTokenInDatabase(Tampung_pengguna[0].ID, Tampung_pengguna[0].Token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			err = tx.Commit(ctx)
			if err != nil {
				panic(err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Berhasil Login!",
			"data":    Tampung_pengguna,
		})
		err = tx.Commit(ctx)
		if err != nil {
			log.Println(err.Error())
		}
		return

	}

}
