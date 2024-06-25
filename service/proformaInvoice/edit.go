package proformaInvoice

import (
	"backend_project_fismed/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func EditPI(c *gin.Context) {
	//  Edit Data PI
	var input service.PerformanceInvoiceDetail
	var response service.PerformanceInvoiceDetail

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	// Proses Inquiry Edit

	response.ID = input.ID
	response.CustomerID = input.CustomerID
	response.Status = input.Status
	response.Divisi = input.Divisi
	response.InvoiceNumber = input.InvoiceNumber
	response.PONumber = input.PONumber
	response.DueDate = input.DueDate
	response.DoctorName = input.DoctorName
	response.PatientName = input.PatientName
	response.TanggalTindakan = input.TanggalTindakan
	response.RM = input.RM
	response.NumberSI = input.NumberSI

	var Subtotal int = 0
	response.ItemDetailPI = make([]service.ResItemDetail, len(input.ItemDetailPI))

	if len(input.ItemDetailPI) > 0 {

		for i, item := range input.ItemDetailPI {
			log.Println("Data Ke :", i)

			response.ItemDetailPI[i].Kat = item.Kat
			response.ItemDetailPI[i].NamaBarang = item.NamaBarang
			response.ItemDetailPI[i].Quantity = item.Quantity
			response.ItemDetailPI[i].HargaSatuan = item.HargaSatuan
			response.ItemDetailPI[i].Discount = item.Discount

			QuantitiInt, err := strconv.Atoi(item.Quantity)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Quantity format"})
				return
			}

			HargaSatuanInt, err := strconv.Atoi(item.HargaSatuan)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Harga Satuan format"})
				return
			}

			DiscountInt, err := strconv.Atoi(item.Discount)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Discount format"})
				return
			}

			subtotalperitem := QuantitiInt*HargaSatuanInt - (QuantitiInt * HargaSatuanInt * DiscountInt / 100)

			Subtotal = Subtotal + subtotalperitem

			subtotalperitemstring := strconv.Itoa(subtotalperitem)

			response.ItemDetailPI[i].SubTotalItem = subtotalperitemstring
			response.ItemDetailPI[i].RPSubTotalItem = "Rp. " + service.FormatRupiah(subtotalperitemstring)
		}
	}

	log.Println("Subtotal :", Subtotal)

	pajak := Subtotal * 11 / 100
	total := pajak + Subtotal

	response.SubTotal = strconv.Itoa(Subtotal)
	response.Pajak = strconv.Itoa(pajak)
	response.Total = strconv.Itoa(total)
	response.TotalRP = "Rp. " + service.FormatRupiah(response.Total)
	response.PajakPPNRP = "Rp. " + service.FormatRupiah(response.Pajak)
	response.SubTotalRP = "Rp. " + service.FormatRupiah(response.SubTotal)

	c.JSON(http.StatusOK, gin.H{"message": "Inquiry Performa Invoice Success !", "data": response, "status": true})
}
