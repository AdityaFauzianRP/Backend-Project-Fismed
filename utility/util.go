package utility

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var DBConnect *pgxpool.Pool

func InitiateDB(db *pgxpool.Pool) {

	DBConnect = db

}

func GenerateToken(userID int) (string, error) {
	privateKey := []byte("secretKey")

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UpdateTokenInDatabase(userID int, newToken string) error {
	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	log.Println("ID :", userID, "New Token", newToken)

	updateQuery := `
        UPDATE users
        SET token = $1
        WHERE id = $2
    `

	_, err = tx.Exec(ctx, updateQuery, newToken, userID)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GenerateNomorInvoice() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "XYZ0123456789"
	length := 10
	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomChar := charset[rand.Intn(len(charset))]
		sb.WriteByte(randomChar)
	}
	return sb.String()
}

func GenerateNomorPO() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABC0123456789"
	length := 10
	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomChar := charset[rand.Intn(len(charset))]
		sb.WriteByte(randomChar)
	}
	return sb.String()
}

func GenerateTigaNomor() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"
	length := 3
	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomChar := charset[rand.Intn(len(charset))]
		sb.WriteByte(randomChar)
	}
	return sb.String()
}

func FormatRupiah(amountStr string) string {
	num, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return ""
	}

	// Format the integer with thousand separators
	formatted := fmt.Sprintf("%d", num)

	// Insert dots every three digits from the end
	var result strings.Builder
	n := len(formatted)

	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteByte('.')
		}
		result.WriteByte(formatted[i])
	}

	return result.String()
}
