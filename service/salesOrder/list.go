package salesOrder

import (
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

func ListDaftar_PO(c *gin.Context) {

	var input model.RequestID
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

	if input.DOK == "PI" {
		query := `
		select 
				COALESCE(a.id, 0) AS id,
				COALESCE(a.customer, '') AS customer,
				COALESCE(a.status, '') AS status,
				COALESCE(a.divisi, '') AS divisi,
				COALESCE(a.invoice_number, '') AS invoice_number,
				COALESCE(a.due_date , '') AS due_date,
				COALESCE(a.doctor_name, '') AS doctor_name,
				COALESCE(a.patient_name, '') AS patient_name,
				COALESCE(a.pajak, '') AS pajak,
				COALESCE(a.total, '') AS total,
				COALESCE(a.sub_total, '') AS sub_total,
				COALESCE(a.tanggal_tindakan, '') AS tanggal_tindakan,
				COALESCE(a.rm, '') AS rm,
				COALESCE(a.number_si, '') AS number_si,
				a.created_at 
			from performance_invoice a where status = 'Diterima'  order by id asc;
	`

		row, err := tx.Query(ctx, query)
		if err != nil {
			log.Println("Gagal Eksekusi Query !")
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query diterima", "status": false})
			return

		}

		defer row.Close()

		var Tampung_data []model.PerformanceInvoiceDetail

		for row.Next() {
			var ambil model.PerformanceInvoiceDetail
			var tanggalAsli time.Time
			err := row.Scan(
				&ambil.ID,
				&ambil.Customer,
				&ambil.Status,
				&ambil.Divisi,
				&ambil.InvoiceNumber,
				&ambil.DueDate,
				&ambil.DoctorName,
				&ambil.PatientName,
				&ambil.PajakPPNRP,
				&ambil.TotalRP,
				&ambil.SubTotalRP,
				&ambil.TanggalTindakan,
				&ambil.RM,
				&ambil.NumberSI,
				&tanggalAsli,
			)

			if err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan Data !", "status": false})
				return
			}

			ambil.PajakPPNRP = "Rp. " + utility.FormatRupiah(ambil.PajakPPNRP)
			ambil.TotalRP = "Rp. " + utility.FormatRupiah(ambil.TotalRP)
			ambil.SubTotalRP = "Rp. " + utility.FormatRupiah(ambil.SubTotalRP)

			ambil.Tanggal = utility.FormatTanggal1(tanggalAsli)

			Tampung_data = append(Tampung_data, ambil)
		}

		if Tampung_data != nil {
			for i, data := range Tampung_data {
				var TampungItem []model.ResItemDetail
				queryOrder := `
				select 
					id , 
					COALESCE("name", '') AS "name",
					COALESCE(quantity, '') AS quantity,
					COALESCE(price, '') AS price,
					COALESCE(discount, '') AS discount,
					COALESCE(sub_total, '') AS sub_total,
					COALESCE(kat, '') AS kat
				from order_items oi where pi_id = $1
			`

				rows, err := tx.Query(ctx, queryOrder, data.ID)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query barang", "status": false})
					return
				}
				defer rows.Close()

				for rows.Next() {
					var ambil model.ResItemDetail

					err := rows.Scan(
						&ambil.Id,
						&ambil.NamaBarang,
						&ambil.Quantity,
						&ambil.HargaSatuan,
						&ambil.Discount,
						&ambil.SubTotalItem,
						&ambil.Kat,
					)
					if err != nil {
						tx.Rollback(ctx)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
						return
					}

					TampungItem = append(TampungItem, ambil)
				}

				Tampung_data[i].ItemDetailPI = TampungItem
			}
		}

		if len(Tampung_data) > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Tampung_data, "status": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
		}
	}

	if input.DOK == "PO" {
		query := `
			select 
				COALESCE(a.id, 0) AS id,
				COALESCE(a.nama_suplier , '') AS customer,
				COALESCE(a.status, '') AS status,
				COALESCE(a.nomor_po , '') AS invoice_number,
				COALESCE(a.pajak, '') AS pajak,
				COALESCE(a.total, '') AS total,
				COALESCE(a.sub_total, '') AS sub_total,
				COALESCE(a.catatan_po, '') AS catatan_po
			from purchase_order a where status = 'DITERIMA'  order by id asc;
	`

		row, err := tx.Query(ctx, query)
		if err != nil {
			log.Println("Gagal Eksekusi Query !")
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return

		}

		defer row.Close()

		var Tampung_data []model.PerformanceInvoiceDetail

		for row.Next() {
			var ambil model.PerformanceInvoiceDetail
			var tanggalAsli time.Time
			err := row.Scan(
				&ambil.ID,
				&ambil.Customer,
				&ambil.Status,
				&ambil.InvoiceNumber,
				&ambil.PajakPPNRP,
				&ambil.TotalRP,
				&ambil.SubTotalRP,
				&ambil.Catatan,
			)

			if err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
				return
			}

			ambil.Tanggal = utility.FormatTanggal1(tanggalAsli)

			Tampung_data = append(Tampung_data, ambil)
		}

		if Tampung_data != nil {
			for i, data := range Tampung_data {
				var TampungItem []model.ResItemDetail
				queryOrder := `
				select 
					id , 
					COALESCE("name", '') AS "name",
					COALESCE(quantity, '') AS quantity,
					COALESCE(price, '') AS price,
					COALESCE(discount, '') AS discount,
					COALESCE(amount , '') AS sub_total,
					COALESCE(kode, '') AS kat
				from item_buyer oi where po_id = $1
			`

				rows, err := tx.Query(ctx, queryOrder, data.ID)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query barang", "status": false})
					return
				}
				defer rows.Close()

				for rows.Next() {
					var ambil model.ResItemDetail

					err := rows.Scan(
						&ambil.Id,
						&ambil.NamaBarang,
						&ambil.Quantity,
						&ambil.HargaSatuan,
						&ambil.Discount,
						&ambil.SubTotalItem,
						&ambil.Kat,
					)
					if err != nil {
						tx.Rollback(ctx)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
						return
					}

					TampungItem = append(TampungItem, ambil)
				}

				Tampung_data[i].ItemDetailPI = TampungItem
			}
		}

		if len(Tampung_data) > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Tampung_data, "status": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
		}
	}

}

func ListDaftar_POFinance(c *gin.Context) {

	var input model.RequestID
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

	if input.DOK == "PO" {
		query := `
			select 
				COALESCE(a.id, 0) AS id,
				COALESCE(a.nama_suplier , '') AS customer,
				COALESCE(a.status, '') AS status,
				COALESCE(a.nomor_po , '') AS invoice_number,
				COALESCE(a.pajak, '') AS pajak,
				COALESCE(a.total, '') AS total,
				COALESCE(a.sub_total, '') AS sub_total,
				COALESCE(a.catatan_po, '') AS catatan_po,
				COALESCE(a.prepared_by, '') AS prepared_by,
				COALESCE(a.prepared_jabatan, '') AS prepared_jabatan,
				COALESCE(a.approved_by, '') AS approved_by,
				COALESCE(a.approved_jabatan, '') AS approved_jabatan,
				COALESCE(a.nomor_si, '') AS surat_jalan
			from purchase_order_copy a where status = 'DITERIMA' order by id asc;
	`

		row, err := tx.Query(ctx, query)
		if err != nil {
			log.Println("Gagal Eksekusi Query !")
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return

		}

		defer row.Close()

		var Tampung_data []model.PerformanceInvoiceDetail

		for row.Next() {
			var ambil model.PerformanceInvoiceDetail
			var tanggalAsli time.Time
			err := row.Scan(
				&ambil.ID,
				&ambil.Customer,
				&ambil.Status,
				&ambil.InvoiceNumber,
				&ambil.PajakPPNRP,
				&ambil.TotalRP,
				&ambil.SubTotalRP,
				&ambil.Catatan,
				&ambil.PreperadBy,
				&ambil.PreperadJabatan,
				&ambil.ApprovedBy,
				&ambil.ApproveJabatan,
				&ambil.NumberSI,
			)

			if err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
				return
			}

			ambil.Nomor_sj = ambil.NumberSI

			ambil.Tanggal = utility.FormatTanggal1(tanggalAsli)

			Tampung_data = append(Tampung_data, ambil)
		}

		if Tampung_data != nil {
			for i, data := range Tampung_data {
				var TampungItem []model.ResItemDetail
				queryOrder := `
				select 
					id , 
					COALESCE("name", '') AS "name",
					COALESCE(quantity, '') AS quantity,
					COALESCE(price, '') AS price,
					COALESCE(discount, '') AS discount,
					COALESCE(amount , '') AS sub_total,
					COALESCE(kode, '') AS kat
				from item_buyer_copy oi where po_id = $1
			`

				rows, err := tx.Query(ctx, queryOrder, data.ID)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query barang", "status": false})
					return
				}
				defer rows.Close()

				for rows.Next() {
					var ambil model.ResItemDetail

					err := rows.Scan(
						&ambil.Id,
						&ambil.NamaBarang,
						&ambil.Quantity,
						&ambil.HargaSatuan,
						&ambil.Discount,
						&ambil.SubTotalItem,
						&ambil.Kat,
					)
					if err != nil {
						tx.Rollback(ctx)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
						return
					}

					TampungItem = append(TampungItem, ambil)
				}

				Tampung_data[i].ItemDetailPI = TampungItem
			}
		}

		if len(Tampung_data) > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Tampung_data, "status": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
		}
	}

}

func ListDaftar_POAdmin(c *gin.Context) {

	var input model.RequestID
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

	if input.DOK == "PI" {
		query := `
		select 
				COALESCE(a.id, 0) AS id,
				COALESCE(a.customer, '') AS customer,
				COALESCE(a.status, '') AS status,
				COALESCE(a.divisi, '') AS divisi,
				COALESCE(a.invoice_number, '') AS invoice_number,
				COALESCE(a.due_date , '') AS due_date,
				COALESCE(a.doctor_name, '') AS doctor_name,
				COALESCE(a.patient_name, '') AS patient_name,
				COALESCE(a.pajak, '') AS pajak,
				COALESCE(a.total, '') AS total,
				COALESCE(a.sub_total, '') AS sub_total,
				COALESCE(a.tanggal_tindakan, '') AS tanggal_tindakan,
				COALESCE(a.rm, '') AS rm,
				COALESCE(a.number_si, '') AS number_si,
				a.created_at 
			from performance_invoice_copy a where status = 'Diterima'  order by id asc;
	`

		row, err := tx.Query(ctx, query)
		if err != nil {
			log.Println("Gagal Eksekusi Query !")
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query diterima", "status": false})
			return

		}

		defer row.Close()

		var Tampung_data []model.PerformanceInvoiceDetail

		for row.Next() {
			var ambil model.PerformanceInvoiceDetail
			var tanggalAsli time.Time
			err := row.Scan(
				&ambil.ID,
				&ambil.Customer,
				&ambil.Status,
				&ambil.Divisi,
				&ambil.InvoiceNumber,
				&ambil.DueDate,
				&ambil.DoctorName,
				&ambil.PatientName,
				&ambil.PajakPPNRP,
				&ambil.TotalRP,
				&ambil.SubTotalRP,
				&ambil.TanggalTindakan,
				&ambil.RM,
				&ambil.NumberSI,
				&tanggalAsli,
			)

			if err != nil {
				tx.Rollback(ctx)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan Data !", "status": false})
				return
			}

			ambil.PajakPPNRP = "Rp. " + utility.FormatRupiah(ambil.PajakPPNRP)
			ambil.TotalRP = "Rp. " + utility.FormatRupiah(ambil.TotalRP)
			ambil.SubTotalRP = "Rp. " + utility.FormatRupiah(ambil.SubTotalRP)
			ambil.Nomor_sj = ambil.NumberSI

			ambil.Tanggal = utility.FormatTanggal1(tanggalAsli)

			Tampung_data = append(Tampung_data, ambil)
		}

		if Tampung_data != nil {
			for i, data := range Tampung_data {
				var TampungItem []model.ResItemDetail
				queryOrder := `
				select 
					id , 
					COALESCE("name", '') AS "name",
					COALESCE(quantity, '') AS quantity,
					COALESCE(price, '') AS price,
					COALESCE(discount, '') AS discount,
					COALESCE(sub_total, '') AS sub_total,
					COALESCE(kat, '') AS kat
				from order_items_copy oi where pi_id = $1
			`

				rows, err := tx.Query(ctx, queryOrder, data.ID)
				if err != nil {
					tx.Rollback(ctx)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query barang", "status": false})
					return
				}
				defer rows.Close()

				for rows.Next() {
					var ambil model.ResItemDetail

					err := rows.Scan(
						&ambil.Id,
						&ambil.NamaBarang,
						&ambil.Quantity,
						&ambil.HargaSatuan,
						&ambil.Discount,
						&ambil.SubTotalItem,
						&ambil.Kat,
					)
					if err != nil {
						tx.Rollback(ctx)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan query", "status": false})
						return
					}

					TampungItem = append(TampungItem, ambil)
				}

				Tampung_data[i].ItemDetailPI = TampungItem
			}
		}

		if len(Tampung_data) > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Data Ditemukan !", "data": Tampung_data, "status": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.StockBarang{}, "status": true})
		}
	}

}

func ListDaftar_POAdmin_edit(c *gin.Context) {

	var input model.PerformanceInvoiceDetail
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

	var totalBaru, subtotalBaru, ppnBaru int

	if len(input.ItemDetailPI) != 0 {
		for _, data := range input.ItemDetailPI {

			log.Println("Nama Barang :", data.NamaBarang)
			log.Println("Id Barang : ", data.Id)

			total, err := utility.CalculateTotal(data.Quantity, data.HargaSatuan, data.Discount)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}

			totalint, _ := strconv.Atoi(total)

			subtotalBaru = subtotalBaru + totalint

			log.Println("Sub Total Baru : ", total)

			totalRp := "Rp. " + utility.FormatRupiah(total)

			query := `UPDATE order_items_copy SET 
                            name = $1, 
                            quantity = $3, 
                            price = $4, 
                            discount = $5, 
                            sub_total = $6 
                        WHERE id = $2`

			_, err = tx.Exec(context.Background(), query, data.NamaBarang, data.Id, data.Quantity, data.HargaSatuan, data.Discount, totalRp)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
		}
	}

	//subtotalBaru

	ppnBaru = subtotalBaru * 11 / 100
	totalBaru = ppnBaru + subtotalBaru

	input.Total = strconv.Itoa(totalBaru)
	input.Pajak = strconv.Itoa(ppnBaru)
	input.SubTotal = strconv.Itoa(subtotalBaru)

	log.Println("Total Baru :", input.Total)
	log.Println("Sub Total Baru :", input.SubTotal)
	log.Println("Pajak Baru :", input.Pajak)

	query := `	
		UPDATE performance_invoice_copy SET 
    	    sub_total = $2, 
    	    pajak = $3, 
    	    total = $4
    	WHERE id = $1
	`

	_, err = tx.Exec(context.Background(), query, input.ID, input.SubTotal, input.Pajak, input.Total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err = tx.Commit(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

}
