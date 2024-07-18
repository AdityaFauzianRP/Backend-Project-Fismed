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

func TanggalSekarang() time.Time {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return time.Time{} // Mengembalikan waktu kosong jika terjadi kesalahan
	}

	// Mengambil waktu saat ini di zona waktu Asia/Jakarta
	now := time.Now().In(location)
	return now
}

// FormatTanggal1 mengembalikan tanggal dalam format tttt - bb - hh
func FormatTanggal1(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTanggal2 mengembalikan tanggal dalam format hh - bb - tttt
func FormatTanggal2(t time.Time) string {
	return t.Format("02-01-2006")
}

// FormatTanggal3 mengembalikan tanggal dalam format tttt - bb - hh  jj - mm - dd
func FormatTanggal3(t time.Time) string {
	return t.Format("2006-01-02 15-04-05")
}

func GenerateToken(userID int, username string) (string, error) {
	privateKey := []byte(username)

	expirationTime := time.Now().Add(8 * time.Hour)

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

func RupiahToNumber(rupiah string) string {
	// Hilangkan "Rp." dan titik-titik
	rupiah = strings.Replace(rupiah, "Rp. ", "", -1)
	rupiah = strings.Replace(rupiah, ".", "", -1)

	// Konversi string ke float64
	fmt.Printf("Angka biasa: %.0f\n", rupiah)
	return rupiah
}

func PersenToNumber(persen string) string {
	// Hilangkan "Rp." dan titik-titik
	persen = strings.Replace(persen, "%", "", -1)

	// Konversi string ke float64
	fmt.Printf("Angka biasa: %.0f\n", persen)
	return persen
}
