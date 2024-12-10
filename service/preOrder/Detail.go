package preOrder

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/service/pengeluaran"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
)

func Detail(c *gin.Context) {
	var input model.RequestID

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		SELECT 
			COALESCE(id, 0) AS id,
			COALESCE(nama_suplier, '') AS nama_suplier,
			COALESCE(nomor_po, '') AS nomor_po,
			COALESCE(tanggal, '') AS tanggal,
			COALESCE(catatan_po, '') AS catatan_po,
			COALESCE(prepared_by, '') AS prepared_by,
			COALESCE(prepared_jabatan, '') AS prepared_jabatan,
			COALESCE(approved_by, '') AS approved_by,
			COALESCE(approved_jabatan, '') AS approved_jabatan,
			COALESCE(status, '') AS status,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total,
			COALESCE(reason, '') AS reason,
			COALESCE(nomor_si, '') AS nomor_si
		FROM purchase_order WHERE id = $1 order by id asc;
	`

	rows, err := tx.Query(ctx, query, input.ID)
	if err != nil {
		tx.Rollback(ctx)
		utility.ResponseError(c, constanta.ErrQuery1)
		return
	}
	defer rows.Close()

	var res model.PurchaseOrder
	for rows.Next() {
		if err := rows.Scan(
			&res.ID,
			&res.NamaSuplier,
			&res.NomorPO,
			&res.Tanggal,
			&res.CatatanPO,
			&res.PreparedBy,
			&res.PreparedJabatan,
			&res.ApprovedBy,
			&res.ApprovedJabatan,
			&res.Status,
			&res.SubTotal,
			&res.Pajak,
			&res.Total,
			&res.Reason,
			&res.NomorSI,
		); err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrScan1)
			return
		}
	}

	QeuryItem := `
		SELECT 
			COALESCE(id, 0) AS id,
			COALESCE(po_id, 0) AS po_id,
			COALESCE(name, '') AS name,
			COALESCE(quantity, '') AS quantity,
			COALESCE(price, '') AS price,
			COALESCE(discount::TEXT, '0') || '' AS discount,
			COALESCE(amount, '') AS amount,
			COALESCE(kode, '') AS kode,
			COALESCE(variable , '') AS variable,
			COALESCE(gudang, '') AS gudang,
			COALESCE(lots, '') AS lots,
			COALESCE(keterangan_barang, '') AS keterangan_baran
		FROM item_buyer WHERE po_id = $1;
	`
	rows2, err := tx.Query(ctx, QeuryItem, res.ID)
	if err != nil {
		tx.Rollback(ctx)
		utility.ResponseError(c, constanta.ErrQuery2)
		return
	}
	defer rows2.Close()

	var resItem []model.ItemBuyer
	for rows2.Next() {
		var get model.ItemBuyer
		if err := rows2.Scan(
			&get.ID,
			&get.POID,
			&get.Name,
			&get.Quantity,
			&get.PriceRP,
			&get.Discount,
			&get.Amount,
			&get.Kode,
			&get.Variable,
			&get.Gudang,
			&get.Lots,
			&get.KeteranganBarang,
		); err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrScan2)
			return
		}

		resItem = append(resItem, get)
	}

	var subtotaltampung int = 0

	if len(resItem) > 0 {
		for i, item := range resItem {
			harga, _ := strconv.Atoi(utility.RupiahToNumber(item.PriceRP))
			quantity, _ := strconv.Atoi(item.Quantity)
			subtotalitem := 0
			subtotalitem = subtotalitem + (harga * quantity)

			item.Amount = strconv.Itoa(subtotalitem)

			subtotaltampung = subtotaltampung + subtotalitem

			resItem[i].Price = utility.RupiahToNumber(item.PriceRP)

		}
	}

	log.Println(res.SubTotal)
	log.Println(res.Total)
	log.Println(res.Pajak)

	res.SubTotalRP = res.SubTotal
	res.TotalRP = res.Total
	res.PajakRP = res.Pajak

	res.SubTotal = utility.RupiahToNumber(res.SubTotal)
	res.Total = utility.RupiahToNumber(res.Total)
	res.Pajak = utility.RupiahToNumber(res.Pajak)

	res.Item = resItem

	if res.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": res, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}

func DetailSO(c *gin.Context) {
	var input model.RequestID

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			return
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			return
		}

	}

	if len(input.ID) >= 4 && input.ID[:4] == "PENG" {
		log.Println("ini dari tambah pengeluaran")
		responseData := pengeluaran.DetailPengeluaran(c, input.ID)

		if responseData.ID != 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data Ditemukan !",
				"data":    responseData,
				"status":  true,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data Tidak Ditemukan !",
				"data":    []model.StockBarang{},
				"status":  true,
			})
		}

		return
	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		SELECT 
			COALESCE(id, 0) AS id,
			COALESCE(nama_suplier, '') AS nama_suplier,
			COALESCE(nomor_po, '') AS nomor_po,
			COALESCE(tanggal, '') AS tanggal,
			COALESCE(catatan_po, '') AS catatan_po,
			COALESCE(prepared_by, '') AS prepared_by,
			COALESCE(prepared_jabatan, '') AS prepared_jabatan,
			COALESCE(approved_by, '') AS approved_by,
			COALESCE(approved_jabatan, '') AS approved_jabatan,
			COALESCE(status, '') AS status,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total,
			COALESCE(reason, '') AS reason,
			COALESCE(nomor_si, '') AS nomor_si
		FROM purchase_order_copy WHERE id = $1 order by id asc;
	`

	rows, err := tx.Query(ctx, query, input.ID)
	if err != nil {
		tx.Rollback(ctx)
		utility.ResponseError(c, constanta.ErrQuery1)
		return
	}
	defer rows.Close()

	var res model.PurchaseOrder
	for rows.Next() {
		if err := rows.Scan(
			&res.ID,
			&res.NamaSuplier,
			&res.NomorPO,
			&res.Tanggal,
			&res.CatatanPO,
			&res.PreparedBy,
			&res.PreparedJabatan,
			&res.ApprovedBy,
			&res.ApprovedJabatan,
			&res.Status,
			&res.SubTotal,
			&res.Pajak,
			&res.Total,
			&res.Reason,
			&res.NomorSI,
		); err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrScan1)
			return
		}
	}

	QeuryItem := `
		SELECT 
			COALESCE(id, 0) AS id,
			COALESCE(po_id, 0) AS po_id,
			COALESCE(name, '') AS name,
			COALESCE(quantity, '') AS quantity,
			COALESCE(price, '') AS price,
			COALESCE(discount::TEXT, '0') || '' AS discount,
			COALESCE(amount, '') AS amount,
			COALESCE(kode, '') AS kode,
			COALESCE(variable , '') AS variable,
			COALESCE(gudang, '') AS gudang,
			COALESCE(lots, '') AS lots
		FROM item_buyer_copy WHERE po_id = $1;
	`
	rows2, err := tx.Query(ctx, QeuryItem, res.ID)
	if err != nil {
		tx.Rollback(ctx)
		utility.ResponseError(c, constanta.ErrQuery2)
		return
	}
	defer rows2.Close()

	var resItem []model.ItemBuyer
	for rows2.Next() {
		var get model.ItemBuyer
		if err := rows2.Scan(
			&get.ID,
			&get.POID,
			&get.Name,
			&get.Quantity,
			&get.PriceRP,
			&get.Discount,
			&get.Amount,
			&get.Kode,
			&get.Variable,
			&get.Gudang,
			&get.Lots,
		); err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrScan2)
			return
		}

		if get.Discount == "" {
			get.Discount = "0"
		}

		resItem = append(resItem, get)
	}

	var subtotaltampung int = 0
	var totaltampung int = 0
	var pajaktampung int = 0

	if len(resItem) > 0 {
		for i, item := range resItem {
			harga, _ := strconv.Atoi(utility.RupiahToNumber(item.PriceRP))
			quantity, _ := strconv.Atoi(item.Quantity)
			subtotalitem := 0
			subtotalitem = subtotalitem + (harga * quantity)

			item.Amount = strconv.Itoa(subtotalitem)

			subtotaltampung = subtotaltampung + subtotalitem

			resItem[i].Price = utility.RupiahToNumber(item.PriceRP)

		}
	}

	var subtotaltampungString string = ""
	var totaltampungString string = ""
	var pajaktampungString string = ""

	pajaktampung = subtotaltampung * 11 / 100
	totaltampung = pajaktampung + subtotaltampung

	subtotaltampungString = strconv.Itoa(subtotaltampung)
	totaltampungString = strconv.Itoa(totaltampung)
	pajaktampungString = strconv.Itoa(pajaktampung)

	res.SubTotalRP = "Rp. " + utility.FormatRupiah(subtotaltampungString)
	res.TotalRP = "Rp. " + utility.FormatRupiah(totaltampungString)
	res.PajakRP = "Rp. " + utility.FormatRupiah(pajaktampungString)

	res.SubTotal = subtotaltampungString
	res.Total = totaltampungString
	res.Pajak = pajaktampungString

	res.Item = resItem

	if res.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": res, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}
