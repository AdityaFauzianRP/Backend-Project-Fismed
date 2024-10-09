package salesOrder

import (
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"time"
)

func ListDaftar_PO(c *gin.Context) {

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

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

		ambil.Tanggal = utility.FormatTanggal1(tanggalAsli)

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
