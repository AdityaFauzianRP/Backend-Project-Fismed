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
)

func Edit_Admin(c *gin.Context) {

	var input model.PurchaseOrder2

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

	if input.NamaSuplier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
		return
	}

	if input.PreparedJabatan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
		return
	}

	if input.PreparedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
		return
	}

	if input.ApprovedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
		return
	}

	if input.ApprovedJabatan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
		return
	}

	for _, item := range input.Item {
		if item.Kode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}

		if item.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}

		if item.Quantity == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}

		if item.Variable == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}

		if item.Lots == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}

		if item.Price == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}

		if item.Gudang == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Lengkap!", "status": false, "RC": 62})
			return
		}
	}

	var subtotal, total, ppn int

	for i, data := range input.Item {

		harga1, err := strconv.Atoi(data.Price)
		if err != nil {
			log.Println("Harga Bukan String !")
			return
		}

		pcs, err := strconv.Atoi(data.Quantity)
		if err != nil {
			log.Println("Quantity Bukan String !")
			return
		}

		disc, err := strconv.Atoi(data.Diskon)
		if err != nil {
			log.Println("Quantity Bukan String !")
			return
		}

		input.Item[i].Amount = "Rp. " + utility.FormatRupiah(strconv.Itoa(harga1*pcs-(harga1*pcs*disc/100)))
		input.Item[i].Price = "Rp. " + utility.FormatRupiah(strconv.Itoa(harga1))

		subtotal = subtotal + harga1*pcs - (harga1 * pcs * disc / 100)
	}

	ppn = subtotal * 11 / 100
	total = ppn + subtotal

	input.Subtotal = "Rp. " + utility.FormatRupiah(strconv.Itoa(subtotal))
	input.Pajak = "Rp. " + utility.FormatRupiah(strconv.Itoa(ppn))
	input.Total = "Rp. " + utility.FormatRupiah(strconv.Itoa(total))

	c.JSON(http.StatusOK, gin.H{"message": "Inquiry Purcase Order Success !", "data": input, "status": true})

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

		queryListRs := `
		SELECT DISTINCT ON (nama_rumah_sakit) 
				COALESCE(nama_rumah_sakit, 'default_name') AS name
				FROM price_list
		ORDER BY nama_rumah_sakit, id ASC;
		`

		rows, err := tx.Query(ctx, queryListRs)
		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
			return
		}
		defer rows.Close()

		var ListRS []model.Customer
		for rows.Next() {
			var res model.Customer
			if err := rows.Scan(
				&res.Name,
			); err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
				return
			}
			ListRS = append(ListRS, res)
		}

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

		// Masukan Ke Table Stok sebagai penambahan Barang
		var cekStok, idGudang int

		for _, data := range input.Item {

			querySelectGudang := `
				select id from gudang g where nama_gudang = $1
			`

			err = tx.QueryRow(ctx, querySelectGudang, data.Gudang).Scan(&idGudang)
			if err != nil {
				tx.Rollback(ctx)
				utility.ResponseError(c, "Error Pengecekan Data Gudang !")
				return
			}

			queryCek := `
				select count(*)  from stock s where nama  = $1 and gudang_id = $2
			`
			err := tx.QueryRow(ctx, queryCek, data.Name, idGudang).Scan(&cekStok)
			if err != nil {
				tx.Rollback(ctx)
				utility.ResponseError(c, "Error Pengecekan Data Stok !")
				return
			}

			if cekStok == 0 {
				log.Println("Tambah Data Baru Di Gudang !")

				queryInsertBarang := `
					INSERT INTO stock (variable, nama, qty, price, gudang_id, kode, lots) 
					VALUES ($1, $2, $3, $4, $5, $6, $7)
				`
				_, err := tx.Exec(ctx, queryInsertBarang, data.Variable, data.Name, data.Quantity, data.Price, idGudang, data.Kode, data.Lots)
				if err != nil {
					tx.Rollback(ctx)
					log.Println("Error Detail : ", err)
					utility.ResponseError(c, "Error Insert Data Barang Baru !")
					return
				}

				if len(ListRS) != 0 {
					for _, rs := range ListRS {

						SCANCEK := 0

						queryCek := `

							select count(*)  from price_list s where nama = $1 and nama_rumah_sakit = $2
						`
						err := tx.QueryRow(ctx, queryCek, data.Name, rs.Name).Scan(&SCANCEK)
						if err != nil {
							tx.Rollback(ctx)
							utility.ResponseError(c, "Error Pengecekan Data Stok !")
							return
						}

						if SCANCEK == 0 {

							var id int
							err = tx.QueryRow(ctx, "SELECT MIN(id) FROM customer WHERE nama_perusahaan = $1;", rs.Name).Scan(&id)
							if err != nil {
								tx.Rollback(ctx)
								c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
								return
							}

							data.Kode = utility.GenerateKode(id, rs.Name)

							queryInsertBarangPerusahaan := `
							INSERT INTO price_list (nama_rumah_sakit, kode, variable, nama, diskon, price, added)
							VALUES ($1, $2, $3, $4, $5, $6, $7)
						`
							_, err = tx.Exec(ctx, queryInsertBarangPerusahaan, rs.Name, data.Kode, data.Variable, data.Name, "0", data.Price, "1")
							if err != nil {
								tx.Rollback(ctx)
								log.Println("Error Detail : ", err)
								utility.ResponseError(c, "Error Insert Data Barang Baru !")
								return
							}
						}

					}
				}

			} else {
				log.Println("Update Data Di Gudang !")

				queryUpdateBarang := `
					UPDATE stock 
					SET variable = $1, nama = $2, qty = qty + $3, price = $4, kode = $5, lots = $8
					WHERE nama = $6 AND gudang_id = $7;
				`
				_, err := tx.Exec(ctx, queryUpdateBarang, data.Variable, data.Name, data.Quantity, data.Price, data.Kode, data.Name, idGudang, data.Lots)
				if err != nil {
					tx.Rollback(ctx)
					log.Println("Error Detail : ", err)
					utility.ResponseError(c, "Error Update Data Barang !")
					return
				}
			}
		}

		QueryInsertDuplicate := `
			INSERT INTO purchase_order_copy
			SELECT * 
			FROM purchase_order
			WHERE id = $1;
		`

		_, err = tx.Exec(context.Background(), QueryInsertDuplicate, input.ID)
		if err != nil {
			tx.Rollback(ctx)
			log.Println("Error Insert Data Duplicate : ", err)
			utility.ResponseError(c, constanta.ErrQuery3)
			return
		}

		QueryInsertDuplicateItemBuyer := `
			INSERT INTO item_buyer_copy
			SELECT * 
			FROM item_buyer
			WHERE po_id = $1;
		`

		_, err = tx.Exec(context.Background(), QueryInsertDuplicateItemBuyer, input.ID)
		if err != nil {
			tx.Rollback(ctx)
			log.Println("QueryInsertDuplicateItemBuyer : ", err)
			utility.ResponseError(c, constanta.ErrQuery3)
			return
		}

		// Masukan Ke Table Pengeluaran sebagai catatan

		QueryPengeluaran := `
				INSERT INTO pengeluaran (
					nama, 
					sub_total,
					pajak,
					total,
					tanggal,
				    status,
				    po_id
				) VALUES ($1, $2, $3, $4, $5, 'PENDING', $6)`

		_, err = tx.Exec(context.Background(), QueryPengeluaran, input.NamaSuplier, input.SubTotal, input.Pajak, input.Total, input.Tanggal, strconv.Itoa(input.ID))
		if err != nil {
			tx.Rollback(ctx)
			log.Println("QueryPengeluaran :", err)
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
	var input model.PurchaseOrder2
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
		input.Nomor_po,
		input.Tanggal,
		input.CatatanPO,
		input.PreparedBy,
		input.PreparedJabatan,
		input.ApprovedBy,
		input.ApprovedJabatan,
		"ADMIN",
		input.Subtotal,
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
						amount,
						kode,
						variable,
						gudang,
					    discount,
					    lots
					) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

				_, err = tx.Exec(context.Background(), QueryItem, input.ID, item.Name, item.Quantity, item.Price, item.Amount, item.Kode, item.Variable, item.Gudang, item.Diskon, item.Lots)
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
						amount=$4,
						kode=$5,
						variable=$6,
						gudang=$7,
						discount=$9,
						lots=$10
					WHERE id=$8
				`

				_, err = tx.Exec(ctx, queryUpdateItem,
					item.Name,
					item.Quantity,
					item.Price,
					item.Amount,
					item.Kode,
					item.Variable,
					item.Gudang,
					item.ID,
					item.Diskon,
					item.Lots,
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

func CancelPO(c *gin.Context) {
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

	QeuryPembatalanPO := `UPDATE purchase_order SET status = 'DIBATALKAN' WHERE id = $1`
	QueryHapuePengeluaran := `DELETE FROM pengeluaran WHERE po_id = $1`
	QueryPembatanPODuplikat := `UPDATE purchase_order_copy SET status = 'DIBATALKAN' WHERE id = $1`

	_, err = tx.Exec(ctx, QeuryPembatalanPO, input.ID)
	if err != nil {
		utility.ResponseError(c, "Error Query Pembatalan PO")
		return
	}

	_, err = tx.Exec(ctx, QueryHapuePengeluaran, strconv.Itoa(input.ID))
	if err != nil {
		utility.ResponseError(c, "Error Query Hapus Pengeluaran")
		return
	}

	_, err = tx.Exec(ctx, QueryPembatanPODuplikat, input.ID)
	if err != nil {
		utility.ResponseError(c, "Error Query Pembatalan PO Duplikat")
		return
	}

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Update Successfully", "status": true})

	c.JSON(http.StatusOK, gin.H{"message": "Data Edit Success !", "status": true})
}
