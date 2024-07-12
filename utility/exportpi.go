package utility

import (
	"backend_project_fismed/model"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"net/http"
)

func ExportPI(c *gin.Context, Input model.PerformanceInvoiceDetail) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(4, 3, 3)       // Mengatur margin kiri, atas, kanan
	pdf.SetAutoPageBreak(true, 3) // Mengatur margin bawah secara otomatis
	pdf.AddPage()

	image := "/9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRPUjogZ2QtanBlZyB2MS4wICh1c2luZyBJSkcgSlBFRyB2NjIpLCBxdWFsaXR5ID0gODIK/9sAQwAGBAQFBAQGBQUFBgYGBwkOCQkICAkSDQ0KDhUSFhYVEhQUFxohHBcYHxkUFB0nHR8iIyUlJRYcKSwoJCshJCUk/9sAQwEGBgYJCAkRCQkRJBgUGCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQkJCQk/8AAEQgAMgAyAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/aAAwDAQACEQMRAD8A+qKDwKR3VFLMwUDkknGKrHVLHH/H3b/9/BTs2tERKcY7s8q8YeNr/UtQmtbK4ktrOFyg8pirS44JJ649qyNI8U6to1ys0F3NIoOWhlcsj+xB/nWbeOpu5zuX/Wv3H941GCG4X5j6Dmvia9XE+1dR3Wp4ksQ5Tumew+F9YTxJrN3qcaSJDDbxQJu6bjln/wDZR+FdXXker+HvFejaFpSaJLNCWLzXflzLGVkbG3cSR8oAx+ea9D0XX7e/toI47hb64UCOeS2UtGsgA3fN0Az719nQoyjh4VJtNy3XY78Pi+abpTVmvx9DZoooqjvPDfip4vvNS1q40eC4ZLC1IjdEOBK+Mnd6gZxjpxmuACLkfKv5V2PxP8PXGjeJrm6MbfZL1/Njl/h3EfMufXIJx6EVy1pZT30yxQRs7McDAzk195g54ajhIzukran5Pmf1qrjZRldu+noZ6wtLN5cabmJIAArtdBfT/B0MepX0YuLr/WW9qOPMbtIx/hQdu5PIFST+HYvBWmLfaoiPfXOfs9m/U+rSeij+6OSSOa4q9vZr+5kuJ5DJJI25mPVjXh+w/tKXPNWpJ6f3n/kenzywWn2/y/4J2OpeINR8VbbnU7ppomO5Ldflij5OPl7n3PNWdD1+98P3aXFpKwUEeZFn5ZF9CP69qxNLH/Evg/3f6mtLT9OudVvI7O0jaSaQ4AA6D1PoB61+a5lia8sbKEW9HZI97CXajNfEz3y2v4Lq3injbKSoHU+oIyKKistMis7OC2UErDGsYJ9AMUV7ydbsfR3mN1XSItYga1ujutnGHj2qd34kHH4VUsPD+ieF7Z5rW0ht1ijJeU8sFAycsea2q5zx+858L3Nrasq3F60dpHu6ZkcIf0JruhHmkot6GVeMYRdVL3kjwnxXr83ifW7m/mJKOSkKn+CIH5R/X6msfYn91fyrovGPg298JX5SVHks3OYbjHDD0PoR6Vk6Vpd5rV7HZafA1xPIcBV6D3J7D3r9Ew9TDqgpQa5Uj8nxNLFPEyjNPmbPZvhz4Y0bUPB2nXF1pttNK6tudkBJ+c12WnaNp+khxY2cFvv5by0Az9ar+F9G/wCEf0Cy03fva3iCs3q3U/qTWrX53XjTlWlUit2z9WwVH2dCEZLVJBRRRUnWFY+vKrXOkBgCPtoOCPSNyKKKqG5lW+A05Y0ljKSIrqRghhkGo7O0t7VStvbxQgnJEaBf5UUVtH+Gzln/ABkWaKKK5zvCiiigD//Z"

	// Mengatur font dan ukuran untuk judul
	pdf.SetFont("Arial", "B", 20)

	// Mendapatkan lebar halaman untuk penempatan tengah
	pageWidth, _ := pdf.GetPageSize()

	// Menghitung posisi X untuk penempatan tengah
	title := "INVOICE"
	titleWidth := pdf.GetStringWidth(title)
	titleX := (pageWidth - titleWidth) / 2

	// Menulis judul di tengah atas
	pdf.SetXY(titleX, 10)
	pdf.CellFormat(titleWidth, 10, title, "", 0, "C", false, 0, "")

	// Menambahkan garis pembatas di bawah judul
	pdf.SetLineWidth(0.5)
	pdf.Line(10, 25, pageWidth-10, 25) // Garis lebih ke bawah

	// Menambahkan tulisan di bawah garis dengan jarak 25 mm
	pdf.SetFont("Arial", "", 11)

	// Mengatur posisi X untuk penulisan teks
	marginLeft := 10.0
	marginRight := 10.0
	columnWidth := (pageWidth - marginLeft - marginRight) / 3

	// Posisi Y awal untuk tulisan di bawah garis
	startY := 35.0

	// SISI KIRI
	pdf.SetXY(marginLeft, startY)
	pdf.CellFormat(columnWidth, 10, "Nomor Invoice", "", 0, "L", false, 0, "")
	pdf.SetXY(columnWidth-20, startY)
	pdf.CellFormat(columnWidth, 10, ":", "", 0, "L", false, 0, "")

	pdf.SetXY(marginLeft, startY+7)
	pdf.CellFormat(columnWidth, 10, "Tanggal Tindakan", "", 0, "L", false, 0, "")
	pdf.SetXY(columnWidth-20, startY+7)
	pdf.CellFormat(columnWidth, 10, ":", "", 0, "L", false, 0, "")

	pdf.SetXY(marginLeft, startY+14)
	pdf.CellFormat(columnWidth, 10, "Tanggal Invoice", "", 0, "L", false, 0, "")
	pdf.SetXY(columnWidth-20, startY+14)
	pdf.CellFormat(columnWidth, 10, ":", "", 0, "L", false, 0, "")

	pdf.SetXY(marginLeft, startY+21)
	pdf.CellFormat(columnWidth, 10, "Jatuh Tempo", "", 0, "L", false, 0, "")
	pdf.SetXY(columnWidth-20, startY+21)
	pdf.CellFormat(columnWidth, 10, ":", "", 0, "L", false, 0, "")

	imgData, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		log.Fatalf("Failed to decode base64 image: %v", err)
	}

	// Create a temporary file to save the decoded image
	tempFile, err := ioutil.TempFile("", "image-*.jpg")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tempFile.Close()

	// Write the decoded image to the temporary file
	_, err = tempFile.Write(imgData)
	if err != nil {
		log.Fatalf("Failed to write image to temporary file: %v", err)
	}

	// Create a new PDF document

	// Add image to PDF
	pdf.Image(tempFile.Name(), 10, 10, 0, 0, false, "", 0, "")
	// SISI TENGAH
	//pdf.SetXY(marginLeft+columnWidth, startY)
	//pdf.CellFormat(columnWidth, 10, "Kelas:", "", 0, "L", false, 0, "")
	//
	//// SISI KANAN
	//pdf.SetXY(marginLeft+2*columnWidth, startY)
	//pdf.CellFormat(columnWidth, 10, "NIM:", "", 0, "L", false, 0, "")

	err = pdf.Output(c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not generate PDF")
	}
}
