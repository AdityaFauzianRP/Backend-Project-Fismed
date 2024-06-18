package service

import "time"

type StockBarang struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Total     string    `json:"total"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty"`
}

type RequestID struct {
	ID string `json:"id"`
}

type PerformanceInvoice struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	ItemName      string `json:"item_name"`
	Discount      string `json:"discount"`
	Status        string `json:"status"`
	Divisi        string `json:"divisi"`
	InvoiceNumber int    `json:"invoice_number"`
	PONumber      int    `json:"po_number"`
	SubTotal      string `json:"sub_total"`
	Pajak         string `json:"pajak"`
	Total         string `json:"total"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdateAt      string `json:"update_at"`
	UpdatedBy     string `json:"updated_by"`
}
