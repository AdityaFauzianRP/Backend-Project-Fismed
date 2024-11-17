package model

type PurchaseOrder struct {
	ID              int          `json:"id,omitempty"`
	NamaSuplier     string       `json:"nama_suplier,omitempty"`
	NomorPO         string       `json:"nomor_po,omitempty"`
	NomorSI         string       `json:"nomor_si,omitempty"`
	Tanggal         string       `json:"tanggal,omitempty"`
	CatatanPO       string       `json:"catatan_po,omitempty"`
	PreparedBy      string       `json:"prepared_by,omitempty"`
	PreparedJabatan string       `json:"prepared_jabatan,omitempty"`
	ApprovedBy      string       `json:"approved_by,omitempty"`
	ApprovedJabatan string       `json:"approved_jabatan,omitempty"`
	SubTotal        string       `json:"sub_total,omitempty"`
	Pajak           string       `json:"pajak,omitempty"`
	Total           string       `json:"total,omitempty"`
	SubTotalRP      string       `json:"sub_total_rp,omitempty"`
	PajakRP         string       `json:"pajak_rp,omitempty"`
	TotalRP         string       `json:"total_rp,omitempty"`
	CreatedAt       string       `json:"created_at,omitempty"`
	CreatedBy       string       `json:"created_by,omitempty"`
	UpdatedAt       string       `json:"updated_at,omitempty"`
	UpdatedBy       string       `json:"updated_by,omitempty"`
	Massage         string       `json:"message,omitempty"`
	Status          string       `json:"status,omitempty"`
	Item            []ItemBuyer  `json:"item,omitempty"`
	ItemDeleted     []ItemBuyer2 `json:"item_deleted,omitempty"`

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
	Kode     string `json:"kode"`
	Variable string `json:"variable"`
	Gudang   string `json:"gudang"`
	Lots     string `json:"lots"`
}

type ItemBuyer2 struct {
	ID       int    `json:"id,omitempty"`
	POID     int    `json:"po_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Price    string `json:"price,omitempty"`
	PriceRP  string `json:"price_rp,omitempty"`
	Discount string `json:"discount,omitempty"`
	Amount   string `json:"amount,omitempty"`
}

type Item struct {
	ID       int64  `json:"id"`
	PoID     int64  `json:"po_id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Variable string `json:"variable"`
	Kode     string `json:"kode"`
	Gudang   string `json:"gudang"`
	Amount   string `json:"amount"`

	Diskon string `json:"discount"`
	Lots   string `json:"lots"`
}

type ItemDeleted struct {
	ID int64 `json:"id"`
}

type PurchaseOrder2 struct {
	ID              int64         `json:"id"`
	NamaSuplier     string        `json:"nama_suplier"`
	CatatanPO       string        `json:"catatan_po"`
	PreparedBy      string        `json:"prepared_by"`
	PreparedJabatan string        `json:"prepared_jabatan"`
	ApprovedBy      string        `json:"approved_by"`
	ApprovedJabatan string        `json:"approved_jabatan"`
	Status          string        `json:"status"`
	Subtotal        string        `json:"sub_total"`
	Pajak           string        `json:"pajak"`
	Total           string        `json:"total"`
	Tanggal         string        `json:"tanggal"`
	Nomor_po        string        `json:"nomor_po"`
	Nomor_si        string        `json:"nomor_si"`
	Item            []Item        `json:"item"`
	ItemDeleted     []ItemDeleted `json:"item_deleted"`
	Reason          string        `json:"reason"`
}
