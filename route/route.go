package route

import (
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"backend_project_fismed/service/pemasukan"
	"backend_project_fismed/service/pengeluaran"
	"backend_project_fismed/service/piutang"
	"backend_project_fismed/service/preOrder"
	"backend_project_fismed/service/proformaInvoice"
	"backend_project_fismed/service/salesOrder"
	"backend_project_fismed/service/stockBarang"
	"log"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	log.Println("Setting up routes")

	router.POST("/api/login", logMiddleware(authentikasi.Login))
	router.POST("/api/token-validate", logMiddleware(authentikasi.TokenValidate))

	// Customer Profilling API
	router.POST("/api/customer-profilling/add", logMiddleware(customerProfilling.Add))
	router.POST("/api/customer-profilling/get-tax-code", logMiddleware(customerProfilling.GetTaxCode))
	router.POST("/api/customer-profilling/get-by-search", logMiddleware(customerProfilling.GetBySearch))

	// Pro forma Invoice API
	router.POST("/api/proforma-invoice/get-all-list", logMiddleware(proformaInvoice.GetAllList))
	router.POST("/api/proforma-invoice/inquiry", logMiddleware(proformaInvoice.Inquiry))
	router.POST("/api/proforma-invoice/posting", logMiddleware(proformaInvoice.Posting))
	router.POST("/api/proforma-invoice/detailPI", logMiddleware(proformaInvoice.DetailPI))
	router.POST("/api/proforma-invoice/editPI-inquiry", logMiddleware(proformaInvoice.EditPI))
	router.POST("/api/proforma-invoice/editPI-posting", logMiddleware(proformaInvoice.PostingEdit_PI))
	router.POST("/api/proforma-invoice/editPI-admin", logMiddleware(proformaInvoice.EditAdmin))
	router.POST("/api/proforma-invoice/divisi-list", logMiddleware(proformaInvoice.DivisiList))
	router.POST("/api/proforma-invoice/rs-list", logMiddleware(proformaInvoice.RumahSakitList))

	// Stock Barang API
	router.POST("/api/stock-barang/add", logMiddleware(stockBarang.Add))
	router.POST("/api/stock-barang/detail", logMiddleware(stockBarang.Detail))
	router.POST("/api/stock-barang/list", logMiddleware(stockBarang.List))
	router.POST("/api/stock-barang/edit", logMiddleware(stockBarang.Edit))
	router.POST("/api/stock-barang/delete", logMiddleware(stockBarang.Delete))

	// Pre Order API
	router.POST("/api/purchase-order/list", logMiddleware(preOrder.ListPO))
	router.POST("/api/purchase-order/detail", logMiddleware(preOrder.Detail))
	router.POST("/api/purchase-order/inquiry", logMiddleware(preOrder.Inquiry))
	router.POST("/api/purchase-order/posting", logMiddleware(preOrder.Posting))
	router.POST("/api/purchase-order/edit/finance", logMiddleware(preOrder.Edit_Finance))
	router.POST("/api/purchase-order/edit/inquiry", logMiddleware(preOrder.Edit_Admin))
	router.POST("/api/purchase-order/edit/posting", logMiddleware(preOrder.Edit_Finance))
	router.POST("/api/purchase-order/edit/delete-item-buyer", logMiddleware(preOrder.DeleteItemBuyer))
	router.POST("/api/purchase-order/edit/posting-edit-admin", logMiddleware(preOrder.Posting_Edit_admin))

	// Sales Order API
	router.POST("/api/sales_order/list", logMiddleware(salesOrder.ListDaftar_PO))

	// Pemasukan API
	router.POST("/api/pemasukan/list", logMiddleware(pemasukan.List))

	// Piutang API
	router.POST("/api/piutang/list", logMiddleware(piutang.List))

	// Hutang API
	router.POST("/api/hutang/list", logMiddleware(piutang.List))

	// Pengeluaran API
	router.POST("/api/pengeluaran/list", logMiddleware(pengeluaran.List))

	log.Println("Routes setup completed")
}

func logMiddleware(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Handler hit: %s %s", c.Request.Method, c.Request.URL.Path)
		handler(c)
	}
}
