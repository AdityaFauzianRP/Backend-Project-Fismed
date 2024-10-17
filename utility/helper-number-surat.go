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

func GenerateINVNumberRAD(invKe int) string {
	divisiSingkatan := "RAD"
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("INV%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}

func GenerateINVNumberOT(invKe int) string {
	divisiSingkatan := "OT"
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("INV%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}

func GenerateINVNumberPO(invKe int) string {
	divisiSingkatan := "PO"
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("INV%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}

func GenerateSJNumberRAD(invKe int) string {
	divisiSingkatan := "RAD"
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("SJ%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}

func GenerateSJNumberOT(invKe int) string {
	divisiSingkatan := "OT"
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("SJ%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}

func GenerateSJNumberPO(invKe int) string {
	divisiSingkatan := "PO"
	currentTime := time.Now()
	bulan := currentTime.Format("01")
	bulanRomawi := constanta.RomanMonths[bulan]
	tahun := currentTime.Format("2006")

	invNumber := fmt.Sprintf("SJ%d/FGI-%s/%s/%s", invKe, divisiSingkatan, bulanRomawi, tahun)

	return invNumber
}
