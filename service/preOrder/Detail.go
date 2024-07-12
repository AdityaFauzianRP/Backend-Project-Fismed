package preOrder

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
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
			COALESCE(total, '') AS total
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
			COALESCE(discount::TEXT, '0') || '%' AS discount,
			COALESCE(amount, '') AS amount
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
			&get.Price,
			&get.Discount,
			&get.Amount,
		); err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrScan2)
			return
		}

		resItem = append(resItem, get)
	}

	res.Item = resItem

	if res.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": res, "status": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
	}

}
