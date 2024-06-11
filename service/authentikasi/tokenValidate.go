package authentikasi

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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

	// ctx := context.Background()
	// tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	// if err != nil {
	// 	panic(err.Error())
	// }

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

		// Hitung berapa lama waktu token telah kedaluwarsa
		currentTime := time.Now().Unix()
		expiredTime := int64(expirationTime)

		timePassed := currentTime - expiredTime
		if timePassed > 600 { // Lebih dari 10 menit (600 detik)
			log.Println("Waktu lebih dari 10 menit setelah kadaluwarsa. Refresh token.")
			// Lakukan proses refresh token di sini

		} else {
			log.Println("Belum lebih dari 10 menit setelah kadaluwarsa. Tidak perlu refresh token.")
		}
		return
	}

	// Token masih dalam waktu yang valid
	log.Println("Token masih valid")
}
