package utility

import (
	"backend_project_fismed/constanta"
	"fmt"
	"time"
)

func GenerateINVNumber(invKe int, divisiSingkatan string) string {
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("INV%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}
