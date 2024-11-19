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

func CalculateTotal(quantityStr, hargaSatuanStr, discountStr string) (string, error) {
	// Konversi string ke angka
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse quantity: %v", err)
	}

	hargaSatuan, err := strconv.ParseFloat(hargaSatuanStr, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse harga_satuan: %v", err)
	}

	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		return "", fmt.Errorf("failed to parse discount: %v", err)
	}

	// Hitung total
	total := float64(quantity) * hargaSatuan * (1 - discount/100)

	totalStr := strconv.FormatFloat(total, 'f', 2, 64)

	totalStr = strings.TrimSuffix(totalStr, ".00")

	return totalStr, nil
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

func ToUpperCase(namaPerusahaan string) string {
	return strings.ToUpper(namaPerusahaan)
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

func CountPIRAD() (error, string) {
	data := ""
	jumlahNoPI := 0

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err, ""
	}
	defer tx.Rollback(ctx)

	countNoPI := `
		UPDATE nomor_penghitungan
		SET jumlah_no_pi = jumlah_no_pi + 1
		RETURNING jumlah_no_pi;
	`

	err = tx.QueryRow(ctx, countNoPI).Scan(&jumlahNoPI)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	data = GenerateINVNumberRAD(jumlahNoPI)

	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Transaction commit failed: %v\n", err)
	}

	return nil, data
}

func CountPIOT() (error, string) {
	data := ""
	jumlahNoPI := 0

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err, ""
	}
	defer tx.Rollback(ctx)

	countNoPI := `
		UPDATE nomor_penghitungan
		SET jumlah_no_pi = jumlah_no_pi + 1
		RETURNING jumlah_no_pi;
	`

	err = tx.QueryRow(ctx, countNoPI).Scan(&jumlahNoPI)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	data = GenerateINVNumberOT(jumlahNoPI)

	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Transaction commit failed: %v\n", err)
	}

	return nil, data
}

func CountPIPO() (error, string) {
	data := ""
	jumlahNoPI := 0

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err, ""
	}
	defer tx.Rollback(ctx)

	countNoPI := `
		UPDATE nomor_penghitungan
		SET jumlah_no_pi = jumlah_no_pi + 1
		RETURNING jumlah_no_pi;
	`

	err = tx.QueryRow(ctx, countNoPI).Scan(&jumlahNoPI)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Transaction commit failed: %v\n", err)
	}

	data = GenerateINVNumberPO(jumlahNoPI)

	return nil, data
}

func CountSJRAD() (error, string) {
	data := ""
	jumlahNoSJ := 0

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err, ""
	}
	defer tx.Rollback(ctx)

	countSJ := `
		UPDATE nomor_penghitungan
		SET jumlah_no_sj = jumlah_no_sj + 1
		RETURNING jumlah_no_sj;
	`

	err = tx.QueryRow(ctx, countSJ).Scan(&jumlahNoSJ)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Transaction commit failed: %v\n", err)
	}

	data = GenerateSJNumberRAD(jumlahNoSJ)

	return nil, data
}

func CountSJOT() (error, string) {
	data := ""
	jumlahNoSJ := 0

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err, ""
	}
	defer tx.Rollback(ctx)

	countSJ := `
		UPDATE nomor_penghitungan
		SET jumlah_no_sj = jumlah_no_sj + 1
		RETURNING jumlah_no_sj;
	`

	err = tx.QueryRow(ctx, countSJ).Scan(&jumlahNoSJ)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	data = GenerateSJNumberOT(jumlahNoSJ)

	return nil, data
}

func CountSJPO() (error, string) {
	data := ""
	jumlahNoSJ := 0

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err, ""
	}
	defer tx.Rollback(ctx)

	countSJ := `
		UPDATE nomor_penghitungan
		SET jumlah_no_sj = jumlah_no_sj + 1
		RETURNING jumlah_no_sj;
	`

	err = tx.QueryRow(ctx, countSJ).Scan(&jumlahNoSJ)

	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	data = GenerateSJNumberPO(jumlahNoSJ)

	return nil, data
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
	rupiah = strings.Replace(rupiah, "Rp. ", "", -1)
	rupiah = strings.Replace(rupiah, ".", "", -1)

	return rupiah
}

func GenerateKode(IDCustomer int, customerName string) string {
	customerPrefix := strings.ToUpper(customerName[:3])

	now := time.Now()
	bulan := now.Format("01")
	tahun := now.Format("2006")
	// Hasilkan kode
	kode := fmt.Sprintf("%02dPL/FGI-RS%s/%s/%s", IDCustomer, customerPrefix, bulan, tahun)
	return kode
}

func PersenToNumber(persen string) string {
	// Hilangkan "Rp." dan titik-titik
	persen = strings.Replace(persen, "%", "", -1)

	// Konversi string ke float64
	fmt.Printf("Angka biasa: %.0f\n", persen)
	return persen
}
