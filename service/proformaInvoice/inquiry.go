package proformaInvoice

import (
	"backend_project_fismed/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Inquiry(c *gin.Context) {
	// Inquiry API

	var input service.ReqInquiryPI
	var response service.ResInquiryPI

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	log.Println("Data Input :", input)

	// Proses Subtotal
	response.Item = make([]service.ResItem, len(input.Item))

	var Subtotal int = 0

	// Proses Subtotal
	if input.Item != nil {
		for i, item := range input.Item {
			log.Println("Item Proses For Inquiry Data:", item)
			response.Item[i].Kat = item.Kat
			response.Item[i].NamaBarang = item.NamaBarang
			response.Item[i].Quantity = item.Quantity
			response.Item[i].HargaSatuan = item.HargaSatuan
			response.Item[i].Discount = item.Discount

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

			response.Item[i].SubTotalItem = subtotalperitemstring
			response.Item[i].SubTotalItemRP = "Rp. " + service.FormatRupiah(subtotalperitemstring)
		}
	}

	pajak := Subtotal * 11 / 100
	total := pajak + Subtotal

	response.SubTotal = strconv.Itoa(Subtotal)
	response.PajakPPN = strconv.Itoa(pajak)
	response.Total = strconv.Itoa(total)
	response.TotalRP = "Rp. " + service.FormatRupiah(response.Total)
	response.PajakPPNRP = "Rp. " + service.FormatRupiah(response.PajakPPN)
	response.SubTotalRP = "Rp. " + service.FormatRupiah(response.SubTotal)
	response.RumahSakit = input.RumahSakit
	response.IdDivisi = input.IdDivisi
	response.Alamat = input.Alamat
	response.IdRumahSakit = input.IdRumahSakit
	response.Tanggal = time.Now().Format("2006-01-02")
	response.NomorPO = "PO/" + service.GenerateTigaNomor() + "/" + service.GenerateNomorPO()
	response.NomorInvoice = "PI/" + service.GenerateTigaNomor() + "/" + service.GenerateNomorInvoice()
	response.NomorSI = "SI/" + service.GenerateTigaNomor() + "/" + service.GenerateNomorInvoice()

	if input.IdDivisi == "1" {
		response.NamaPasien = ""
		response.NamaDokter = ""
		response.RM = ""

	} else if input.IdDivisi == "2" {
		response.NamaPasien = input.NamaPasien
		response.NamaDokter = "DR. " + input.NamaDokter
		response.RM = input.RM
		response.TanggalTindakan = input.TanggalTindakan
	}

	ExpireDate, err := strconv.Atoi(input.JatuhTempo)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Discount format"})
		return
	}

	response.JatuhTempo = time.Now().AddDate(0, 0, ExpireDate).Format("2006-01-02")

	// Return response

	c.JSON(http.StatusOK, gin.H{"message": "Inquiry Performa Invoice Success !", "data": response, "status": true})

}
