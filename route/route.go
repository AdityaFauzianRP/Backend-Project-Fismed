package route

import (
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"backend_project_fismed/service/gudang"
	"backend_project_fismed/service/pemasukan"
	"backend_project_fismed/service/pengeluaran"
	"backend_project_fismed/service/piutang"
	"backend_project_fismed/service/preOrder"
	price_list "backend_project_fismed/service/price-list"
	"backend_project_fismed/service/proformaInvoice"
	"backend_project_fismed/service/salesOrder"
	"backend_project_fismed/service/stockBarang"
	"backend_project_fismed/service/stok"
	"github.com/gin-gonic/gin"
	"log"
)

func Routes(router *gin.Engine) {
	log.Println("Setting up routes")

	router.POST("/api/login", logMiddleware(authentikasi.Login))
	router.POST("/api/token-validate", logMiddleware(authentikasi.TokenValidate))

	//  Stok Baru APIL
	router.POST("/api/stok/listbygudang", logMiddleware(stok.ListByGudangId))
	router.POST("/api/stok/list-customer", logMiddleware(stok.ListBarang))
	router.POST("/api/stok/list-proses", logMiddleware(stok.ListBarangProses))

	router.POST("/api/barang-terjual", logMiddleware(stok.ListBarangTerjual))

	//  Gudang API
	router.POST("/api/gudang/list", logMiddleware(gudang.GudangListData))
	router.POST("/api/gudang/tambah", logMiddleware(gudang.TambahGudang))
	router.POST("/api/gudang/delete", logMiddleware(gudang.HapusGudang))

	//  Akun API
	router.POST("/api/akun/list", logMiddleware(gudang.AkunListData))
	router.POST("/api/akun/tambah", logMiddleware(gudang.TambahAkun))
	router.POST("/api/akun/delete", logMiddleware(gudang.HapusAkun))

	//  PRICE Baru APIL
	router.POST("/api/price/list", logMiddleware(price_list.PriceList))
	router.POST("/api/price/ListByCustomer", logMiddleware(price_list.ListByCustomer))
	router.POST("/api/price/SetPrice", logMiddleware(price_list.SetPrice))

	// Customer Profilling API
	router.POST("/api/profile/ListByCustomer", logMiddleware(customerProfilling.ListProfile))
	router.POST("/api/profile/DetailProfile", logMiddleware(customerProfilling.DetailProfile))
	router.POST("/api/profile/EditDetailProfile", logMiddleware(customerProfilling.EditDetailProfile))

	router.POST("/api/customer-profilling/add", logMiddleware(customerProfilling.Add))
	router.POST("/api/customer-profilling/get-tax-code", logMiddleware(customerProfilling.GetTaxCode))
	router.POST("/api/customer-profilling/get-by-search", logMiddleware(customerProfilling.GetBySearch))
	router.POST("/api/proforma-invoice/rs-list", logMiddleware(customerProfilling.RumahSakitList))
	router.POST("/api/proforma-invoice/supplier", logMiddleware(customerProfilling.SupplierList))
	router.POST("/api/proforma-invoice/dr-list", logMiddleware(customerProfilling.DokterList))
	router.POST("/api/proforma-invoice/dr-listn", logMiddleware(customerProfilling.DokterListbyname))

	router.POST("/api/proforma-invoice/rs-lists", logMiddleware(customerProfilling.RumahSakitListS))
	router.POST("/api/proforma-invoice/rs-listc", logMiddleware(customerProfilling.RumahSakitListC))
	// Pro forma Invoice API
	router.POST("/api/proforma-invoice/get-all-list", logMiddleware(proformaInvoice.GetAllList))
	router.POST("/api/proforma-invoice/get-all-list-so", logMiddleware(proformaInvoice.GetAllListSO))
	router.POST("/api/proforma-invoice/inquiry", logMiddleware(proformaInvoice.InquiryPI))
	router.POST("/api/proforma-invoice/posting", logMiddleware(proformaInvoice.PostingPI))
	router.POST("/api/proforma-invoice/detailPI", logMiddleware(proformaInvoice.DetailPI))
	router.POST("/api/proforma-invoice/detailPI-so", logMiddleware(proformaInvoice.DetailPISO))
	router.POST("/api/proforma-invoice/detailPI-so-logistik", logMiddleware(proformaInvoice.DetailPISOLogistik))
	router.POST("/api/proforma-invoice/editPI-inquiry", logMiddleware(proformaInvoice.EditPI))
	router.POST("/api/proforma-invoice/editPI-posting", logMiddleware(proformaInvoice.PostingEdit_PI))
	router.POST("/api/proforma-invoice/editPI-admin", logMiddleware(proformaInvoice.EditAdmin))
	router.POST("/api/proforma-invoice/divisi-list", logMiddleware(proformaInvoice.DivisiList))

	router.POST("/api/proforma-invoice/cancel", logMiddleware(proformaInvoice.CancelPI))

	// Stock Barang API
	router.POST("/api/stock-barang/add", logMiddleware(stockBarang.Add))
	router.POST("/api/stock-barang/detail", logMiddleware(stockBarang.Detail))
	router.POST("/api/stock-barang/list", logMiddleware(stockBarang.List))
	router.POST("/api/stock-barang/edit", logMiddleware(stockBarang.Edit))
	router.POST("/api/stock-barang/delete", logMiddleware(stockBarang.Delete))

	// Pre Order API
	router.POST("/api/purchase-order/cancel", logMiddleware(preOrder.CancelPO))
	router.POST("/api/purchase-order/list", logMiddleware(preOrder.ListPO))
	router.POST("/api/purchase-order/detail", logMiddleware(preOrder.Detail))
	router.POST("/api/purchase-order/list-so", logMiddleware(preOrder.ListPOSO))
	router.POST("/api/purchase-order/detail-so", logMiddleware(preOrder.DetailSO))
	router.POST("/api/purchase-order/inquiry", logMiddleware(preOrder.InquiryPO))
	router.POST("/api/purchase-order/posting", logMiddleware(preOrder.Posting))
	router.POST("/api/purchase-order/edit/finance", logMiddleware(preOrder.Edit_Finance))
	router.POST("/api/purchase-order/edit/inquiry", logMiddleware(preOrder.Edit_Admin))
	router.POST("/api/purchase-order/edit/posting", logMiddleware(preOrder.Edit_Finance))
	router.POST("/api/purchase-order/edit/delete-item-buyer", logMiddleware(preOrder.DeleteItemBuyer))
	router.POST("/api/purchase-order/edit/posting-edit-admin", logMiddleware(preOrder.Posting_Edit_admin))

	// Sales Order API
	router.POST("/api/sales_order/list", logMiddleware(salesOrder.ListDaftar_PO))
	router.POST("/api/sales_order/list/finance", logMiddleware(salesOrder.ListDaftar_POFinance))
	router.POST("/api/sales_order/list/admin", logMiddleware(salesOrder.ListDaftar_POAdmin))
	router.POST("/api/sales_order/list/admin/edit/nama-barang-pi-so", logMiddleware(salesOrder.ListDaftar_POAdmin_edit))

	// Pemasukan API
	router.POST("/api/pemasukan/list", logMiddleware(pemasukan.List))

	// Piutang API
	router.POST("/api/piutang/list", logMiddleware(piutang.ListPiutang))
	router.POST("/api/piutang/lunas", logMiddleware(piutang.LunasPiutang))

	// Hutang API
	router.POST("/api/hutang/list", logMiddleware(piutang.List))
	router.POST("/api/hutang/lunas", logMiddleware(piutang.Lunas))

	// Pengeluaran API
	router.POST("/api/pengeluaran/list", logMiddleware(pengeluaran.List))
	router.POST("/api/pengeluaran/inquiry", logMiddleware(pengeluaran.Inquiry))
	router.POST("/api/pengeluaran/posting", logMiddleware(pengeluaran.Posting))

	router.GET("/api/check", logMiddleware(authentikasi.HealthCheckHandler))

	log.Println("Routes setup completed")
}

func logMiddleware(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Handler hit: %s %s", c.Request.Method, c.Request.URL.Path)
		handler(c)
	}
}
