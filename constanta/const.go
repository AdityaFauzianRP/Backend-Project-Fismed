package constanta

const (
	Ortopedi  string = "Ortopedi"
	Radiologi string = "Radiologi"

	//	Status Error
	ErrQuery1 string = "Gagal Eksekusi Query 1"
	ErrQuery2 string = "Gagal Eksekusi Query 2"
	ErrQuery3 string = "Gagal Eksekusi Query 3"
	ErrQuery4 string = "Gagal Eksekusi Query 4"
	ErrQuery5 string = "Gagal Eksekusi Query 5"
	ErrQuery6 string = "Gagal Eksekusi Query 6"

	//  Scan Error Massage
	ErrScan1 string = "Gagal Eksekusi Scan 1"
	ErrScan2 string = "Gagal Eksekusi Scan 2"
	ErrScan3 string = "Gagal Eksekusi Scan 3"
	ErrScan4 string = "Gagal Eksekusi Scan 4"
	ErrScan5 string = "Gagal Eksekusi Scan 5"
	ErrScan6 string = "Gagal Eksekusi Scan 6"

	ErrCommit string = "Gagal Commit Ke Database"

	Diterima string = "DITERIMA"
	Ditolak  string = "DIPROSES"
	Diproses string = "DITOLAK"

	CustomerRS         string = "1"
	CustomerNonRS      string = "2"
	CustomerAsSupplier string = "3"
)

var RomanMonths = map[string]string{
	"01": "I",
	"02": "II",
	"03": "III",
	"04": "IV",
	"05": "V",
	"06": "VI",
	"07": "VII",
	"08": "VIII",
	"09": "IX",
	"10": "X",
	"11": "XI",
	"12": "XII",
}
