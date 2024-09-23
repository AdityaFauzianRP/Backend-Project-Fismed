package authentikasi

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func TokenValidate(c *gin.Context) {
	type Request struct {
		Token string `json:"token"`
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
	defer tx.Rollback(ctx)

	QueryGetByToke := `
		select u.id, u.username from users u where u.token = $1;
	`

	var ID int
	var USERNAME string

	err = tx.QueryRow(ctx, QueryGetByToke, input.Token).Scan(&ID, &USERNAME)
	if err != nil {
		log.Println("Error parsing token:", err)
		c.JSON(400, gin.H{
			"message": "User tidak valid !",
			"status":  false,
		})
		return
	}

	tokenString := input.Token

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Println("Error parsing token:", err)
		return
	}

	// Periksa informasi dalam token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Invalid token claims")
		return
	}

	// Periksa waktu kadaluarsa token
	expirationTime := claims["exp"].(float64)
	if expirationTime < float64(jwt.TimeFunc().Unix()) {
		log.Println("Token telah kedaluwarsa")

		query := `
			UPDATE users 
			SET "token" = $2
			WHERE id = $1;
			`

		_, err = tx.Exec(context.Background(), query, ID, "Empty")
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}

		if err := tx.Commit(ctx); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Token Is Not Active !",
			"status":  false,
		})

		return
	}

	// Token masih dalam waktu yang valid
	log.Println("Token masih valid")

	c.JSON(http.StatusOK, gin.H{
		"message": "Token Is Active !",
		"status":  true,
	})
}
