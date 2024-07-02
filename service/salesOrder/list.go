package salesOrder

import (
	"backend_project_fismed/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func ListDaftar_PO(c *gin.Context) {
	//	List Daftar PO di ambil dari PI yang di terima

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}

	query := `
		select 
			COALESCE(id, 0) AS id,
			COALESCE(customer_id, 0) AS customer_id,
			COALESCE(status, '') AS status,
			COALESCE(divisi, '') AS divisi,
			COALESCE(invoice_number, '') AS invoice_number,
			COALESCE(po_number, '') AS po_number,
			COALESCE(due_date , '') AS due_date,
			COALESCE(doctor_name, '') AS doctor_name,
			COALESCE(patient_name, '') AS patient_name,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total,
			COALESCE(sub_total, '') AS sub_total,
			COALESCE(tanggal_tindakan, '') AS tanggal_tindakan,
			COALESCE(rm, '') AS rm,
			COALESCE(number_si, '') AS number_si
		from performance_invoice pi2 where status = 'Diterima' order by id asc;
	`

	row, err := tx.Query(ctx, query)
	if err != nil {
		log.Println("Gagal Eksekusi Query !")
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return

	}

	defer row.Close()

	var Tampung_data []model.PerformanceInvoiceDetail

	for row.Next() {
		var ambil model.PerformanceInvoiceDetail
		err := row.Scan(
			&ambil.ID,
			&ambil.CustomerID,
			&ambil.Status,
			&ambil.Divisi,
			&ambil.InvoiceNumber,
			&ambil.PONumber,
			&ambil.DueDate,
			&ambil.DoctorName,
			&ambil.PatientName,
			&ambil.PajakPPNRP,
			&ambil.TotalRP,
			&ambil.SubTotalRP,
			&ambil.TanggalTindakan,
			&ambil.RM,
			&ambil.NumberSI,
		)

		if err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Scan Data !", "status": false})
			return
		}

		Tampung_data = append(Tampung_data, ambil)
	}

	var TampungItem []model.ResItemDetail

	if Tampung_data != nil {
		for i, data := range Tampung_data {
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
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
