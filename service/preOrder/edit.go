package preOrder

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Edit_Admin(c *gin.Context) {

	var input model.PurchaseOrder
	var res model.PurchaseOrder

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	}

	log.Println("Data Input :", input)

	res.Item = make([]model.ItemBuyer, len(input.Item))
	res.ItemDeleted = input.ItemDeleted
	res.ID = input.ID

	var Subtotal int

	if len(input.Item) > 0 {
		for i, item := range input.Item {
			log.Println("Item Barang Ke :", i)
			res.Item[i].ID = item.ID
			res.Item[i].POID = item.POID
			res.Item[i].Name = item.Name
			res.Item[i].Quantity = item.Quantity
			res.Item[i].Price = "Rp. " + utility.FormatRupiah(item.Price)

			QuantitiInt, err := strconv.Atoi(item.Quantity)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Quantity format"})
				return
			}

			item.Price = utility.RupiahToNumber(item.Price)

			HargaSatuanInt, err := strconv.Atoi(item.Price)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Harga Satuan format"})
				return
			}

			item.Discount = utility.PersenToNumber(item.Discount)

			DiscountInt, err := strconv.Atoi(item.Discount)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid Discount format"})
				return
			}

			res.Item[i].Price = "Rp. " + utility.FormatRupiah(item.Price)
			res.Item[i].Discount = item.Discount + "%"

			subtotalperitem := QuantitiInt*HargaSatuanInt - (QuantitiInt * HargaSatuanInt * DiscountInt / 100)

			Subtotal = Subtotal + subtotalperitem

			subtotalperitemstring := strconv.Itoa(subtotalperitem)

			res.Item[i].Amount = "Rp. " + utility.FormatRupiah(subtotalperitemstring)
		}

		pajak := Subtotal * 11 / 100
		total := pajak + Subtotal

		res.Total = "Rp. " + utility.FormatRupiah(strconv.Itoa(total))
		res.Pajak = "Rp. " + utility.FormatRupiah(strconv.Itoa(pajak))
		res.SubTotal = "Rp. " + utility.FormatRupiah(strconv.Itoa(Subtotal))

		currentTime := time.Now()
		timeString := currentTime.Format("2006-01-02")

		res.NamaSuplier = input.NamaSuplier
		res.NomorPO = utility.GenerateNomorPO()
		res.Tanggal = timeString
		res.CatatanPO = input.CatatanPO
		res.PreparedBy = input.PreparedBy
		res.PreparedJabatan = input.PreparedJabatan
		res.ApprovedBy = input.ApprovedBy
		res.ApprovedJabatan = input.ApprovedJabatan

		c.JSON(http.StatusOK, gin.H{"message": "Inquiry Purcase Order Success !", "data": res, "status": true})

	} else {
		utility.ResponseError(c, "Item Empty!")
	}

}

func Edit_Finance(c *gin.Context) {
	var input model.PurchaseOrder
	//var response model.PurchaseOrder

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	}

	log.Println("Data Input :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	if input.Status == "DITOLAK" {
		query := `
			UPDATE purchase_order
			SET
				status = $1,
				reason = $2,
				updated_at = now(),
				updated_by = 'admin'
			WHERE id = $3;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Status,
			input.Reason,
			input.ID,
		)

		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery1)
			return
		}

		if err := tx.Commit(ctx); err != nil {
			utility.ResponseError(c, constanta.ErrCommit)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data Edit Success !", "status": true})
	}

	if input.Status == "DITERIMA" {
		query := `
			UPDATE purchase_order
			SET
				status = $1,
				updated_at = now(),
				updated_by = 'admin'
			WHERE id = $2;
			`

		_, err = tx.Exec(context.Background(), query,
			input.Status,
			input.ID,
		)

		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery2)
			return
		}

		// Masukan Ke Table Pengeluaran sebagai catatan

		QueryPengeluaran := `
				INSERT INTO pengeluaran (
					nama, 
					sub_total,
					pajak,
					total,
					tanggal  
				) VALUES ($1, $2, $3, $4, $5)`

		_, err = tx.Exec(context.Background(), QueryPengeluaran, input.NamaSuplier, input.SubTotal, input.Pajak, input.Total, input.Tanggal)
		if err != nil {
			tx.Rollback(ctx)
			utility.ResponseError(c, constanta.ErrQuery3)
			return
		}

		if err := tx.Commit(ctx); err != nil {
			utility.ResponseError(c, constanta.ErrCommit)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data Edit Success !", "status": true})
	}
}

func Posting_Edit_admin(c *gin.Context) {
	var input model.PurchaseOrder
	//var res model.PurchaseOrder

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	}

	log.Println("Data Input :", input)

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE public.purchase_order
		SET 
		    nama_suplier=$1, 
		    nomor_po=$2, 
		    tanggal=$3, 
		    catatan_po=$4, 
		    prepared_by=$5, 
		    prepared_jabatan=$6, 
		    approved_by=$7, 
		    approved_jabatan=$8, 
		    status='DIPROSES',
		    updated_at=now(), 
		    updated_by=$9, 
		    sub_total=$10, 
		    pajak=$11, 
		    total=$12, 
		    reason=''
		WHERE id=$13
	`

	_, err = tx.Exec(ctx, query,
		input.NamaSuplier,
		input.NomorPO,
		input.Tanggal,
		input.CatatanPO,
		input.PreparedBy,
		input.PreparedJabatan,
		input.ApprovedBy,
		input.ApprovedJabatan,
		utility.FormatTanggal1(time.Now()), // Ensure this function returns a suitable format
		input.SubTotal,
		input.Pajak,
		input.Total,
		input.ID,
	)

	if err != nil {
		utility.ResponseError(c, "Error Qeury Update On Table PO")
		return
	}

	if len(input.Item) > 0 {
		for _, item := range input.Item {
			if item.ID == 0 {
				QueryItem := `
				INSERT INTO item_buyer (
					po_id, 
					name, 
					quantity, 
					price, 
					discount, 
					amount
				) VALUES ($1, $2, $3, $4, $5, $6)`

				_, err = tx.Exec(context.Background(), QueryItem, input.ID, item.Name, item.Quantity, item.Price, item.Discount, item.Amount)
				if err != nil {
					tx.Rollback(ctx)
					utility.ResponseError(c, "Error on Add New Item")
					return
				}

				log.Println("Data Item Baru : ", item.Name, " Berhasil di Input!")
			} else {
				queryUpdateItem := `
					UPDATE public.item_buyer
					SET 
						name=$1, 
						quantity=$2, 
						price=$3, 
						discount=$4, 
						amount=$5
					WHERE id=$6
				`

				_, err = tx.Exec(ctx, queryUpdateItem,
					item.Name,
					item.Quantity,
					item.Price,
					item.Discount,
					item.Amount,
					item.ID,
				)

				if err != nil {
					utility.ResponseError(c, "Error Qeury Update On Table ITEM on Condition Update")
					return
				}
			}
		}
	}

	if len(input.ItemDeleted) > 0 {
		for _, hapus := range input.ItemDeleted {
			if hapus.ID != 0 {
				queryDeleter := "DELETE FROM item_buyer WHERE id = $1"
				_, err = tx.Exec(ctx, queryDeleter, hapus.ID)

				if err != nil {
					utility.ResponseError(c, "Error Query Delete Item By Id")
					return
				}
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		utility.ResponseError(c, "Error Commit Transaction Database!")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Update successfully", "status": true})
}
