package model

type PurchaseOrder struct {
	ID              int         `json:"id,omitempty"`
	NamaSuplier     string      `json:"nama_suplier,omitempty"`
	NomorPO         string      `json:"nomor_po,omitempty"`
	Tanggal         string      `json:"tanggal,omitempty"`
	CatatanPO       string      `json:"catatan_po,omitempty"`
	PreparedBy      string      `json:"prepared_by,omitempty"`
	PreparedJabatan string      `json:"prepared_jabatan,omitempty"`
	ApprovedBy      string      `json:"approved_by,omitempty"`
	ApprovedJabatan string      `json:"approved_jabatan,omitempty"`
	SubTotal        string      `json:"sub_total,omitempty"`
	Pajak           string      `json:"pajak,omitempty"`
	Total           string      `json:"total,omitempty"`
	SubTotalRP      string      `json:"sub_total_rp,omitempty"`
	PajakRP         string      `json:"pajak_rp,omitempty"`
	TotalRP         string      `json:"total_rp,omitempty"`
	CreatedAt       string      `json:"created_at,omitempty"`
	CreatedBy       string      `json:"created_by,omitempty"`
	UpdatedAt       string      `json:"updated_at,omitempty"`
	UpdatedBy       string      `json:"updated_by,omitempty"`
	Massage         string      `json:"message,omitempty"`
	Status          string      `json:"status,omitempty"`
	Item            []ItemBuyer `json:"item,omitempty"`
	ItemDeleted     []ItemBuyer `json:"item_deleted,omitempty"`

	Reason string `json:"reason,omitempty"`
}

// ItemBuyer represents the item_buyer table
type ItemBuyer struct {
	ID       int    `json:"id,omitempty"`
	POID     int    `json:"po_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Price    string `json:"price,omitempty"`
	PriceRP  string `json:"price_rp,omitempty"`
	Discount string `json:"discount,omitempty"`
	Amount   string `json:"amount,omitempty"`
}
