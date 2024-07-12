package route

import (
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"backend_project_fismed/service/pemasukan"
	"backend_project_fismed/service/piutang"
	"backend_project_fismed/service/preOrder"
	"backend_project_fismed/service/proformaInvoice"
	"backend_project_fismed/service/salesOrder"
	"backend_project_fismed/service/stockBarang"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/api/login", authentikasi.Login)
	router.POST("/api/token-validate", authentikasi.TokenValidate)
	//router.POST("/api/image", authentikasi.Image)

	//	Customer Profilling API Start

	router.POST("/api/customer-profilling/add", customerProfilling.Add)
	router.POST("/api/customer-profilling/get-tax-code", customerProfilling.GetTaxCode)
	router.POST("/api/customer-profilling/get-by-search", customerProfilling.GetBySearch)

	//	Customer Profilling API End

	//	Pro forma Invoice API Start

	router.POST("/api/proforma-invoice/get-all-list", proformaInvoice.GetAllList)
	router.POST("/api/proforma-invoice/inquiry", proformaInvoice.Inquiry)
	router.POST("/api/proforma-invoice/posting", proformaInvoice.Posting)
	router.POST("/api/proforma-invoice/detailPI", proformaInvoice.DetailPI)
	router.POST("/api/proforma-invoice/editPI-inquiry", proformaInvoice.EditPI)
	router.POST("/api/proforma-invoice/editPI-posting", proformaInvoice.PostingEdit_PI)
	router.POST("/api/proforma-invoice/editPI-admin", proformaInvoice.EditAdmin)
	router.POST("/api/proforma-invoice/divisi-list", proformaInvoice.DivisiList)
	router.POST("/api/proforma-invoice/rs-list", proformaInvoice.RumahSakitList)

	//	Pro forma Invoice API End

	//	Stock Barang API Start

	router.POST("/api/stock-barang/add", stockBarang.Add)
	router.POST("/api/stock-barang/detail", stockBarang.Detail)
	router.POST("/api/stock-barang/list", stockBarang.List)
	router.POST("/api/stock-barang/edit", stockBarang.Edit)
	router.POST("/api/stock-barang/delete", stockBarang.Delete)

	//	Stock Barang API End

	//	Pre Order API Start

	router.POST("/api/purchase-order/list", preOrder.ListPO)
	router.POST("/api/purchase-order/detail", preOrder.Detail)
	router.POST("/api/purchase-order/inquiry", preOrder.Inquiry)
	router.POST("/api/purchase-order/posting", preOrder.Posting)
	router.POST("/api/purchase-order/edit/finance", preOrder.Edit_Finance)
	router.POST("/api/purchase-order/edit/inquiry", preOrder.Edit_Admin)
	router.POST("/api/purchase-order/edit/posting", preOrder.Edit_Finance)

	//  Pre Order API End

	//	Sales Order API Start

	router.POST("/api/sales_order/list", salesOrder.ListDaftar_PO)

	//  Sales Order API End

	//	Pemasukan API Start

	router.POST("/api/pemasukan/list", pemasukan.List)

	//  Pemasukan API End

	//	Piutang API Start

	router.POST("/api/piutang/list", piutang.List)

	//  Piutang API End
}
