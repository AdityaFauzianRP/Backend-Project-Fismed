package pengeluaran

import (
	"backend_project_fismed/constanta"
	"backend_project_fismed/model"
	"backend_project_fismed/utility"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func List(c *gin.Context) {

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	query := `
		select 
			COALESCE(id, 0) AS id,
			COALESCE(nama, '') AS nama,
			COALESCE(sub_total , '') AS sub_total,
			COALESCE(pajak, '') AS pajak,
			COALESCE(total, '') AS total,
			tanggal,
			COALESCE(po_id, '') AS po_id
		from pengeluaran p where status = 'DITERIMA' order by id desc 
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "status": false})
		return
	}
	defer rows.Close()

	var Responses []model.Pemasukan
	var total int
	for rows.Next() {
		var res model.Pemasukan
		if err := rows.Scan(
			&res.Id,
			&res.Nama,
			&res.Nominal,
			&res.Pajak,
			&res.Amount,
			&res.Tanggal,
			&res.Piid,
		); err != nil {
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
			return
		}
		Amount, _ := strconv.Atoi(res.Amount)
		total = total + Amount

		res.Nominal = "Rp. " + utility.FormatRupiah(res.Nominal)
		res.Pajak = "Rp. " + utility.FormatRupiah(res.Pajak)
		res.Amount = "Rp. " + utility.FormatRupiah(res.Amount)

		Responses = append(Responses, res)
	}

	total_pemasukan := utility.FormatRupiah(strconv.Itoa(total))

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows", "status": false})
		return
	}

	if len(Responses) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data Ditemukan !",
			"data":    Responses,
			"total":   "Rp. " + total_pemasukan,
			"status":  true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "data": []model.Customer{}, "status": true})
	}
}

func DetailPengeluaran(c *gin.Context, PoId string) model.PurchaseOrder {
	var res model.PurchaseOrder
	//var resItem []model.ItemBuyer
	var freeData string

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	Query1 := `
		SELECT free_data FROM pengeluaran where po_id = $1
	`

	err = tx.QueryRow(ctx, Query1, PoId).Scan(&freeData)

	if err != nil {
		log.Println("Error query detail pengeluaran : ", err)
		return model.PurchaseOrder{}
	}
	log.Println("Free Data : ", freeData)

	err = json.Unmarshal([]byte(freeData), &res)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return model.PurchaseOrder{}
	}

	res.ID = 1
	res.NomorPO = "-"
	res.NomorSI = "-"

	for i, data := range res.Item {
		if data.Discount == "" {
			res.Item[i].Discount = "0"
			res.Item[i].Gudang = "-"
			res.Item[i].Kode = "-"
			res.Item[i].Variable = "-"
			res.Item[i].Lots = "-"
		}
	}

	res.TotalRP = res.Total
	res.SubTotalRP = res.SubTotal
	res.PajakRP = res.Pajak

	log.Println("Response : res", res)
	return res

}

func Inquiry(c *gin.Context) {
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
	var subtotal, total, ppn int
	input.Tanggal = utility.FormatTanggal2(time.Now())

	for i, data := range input.Item {

		harga1, err := strconv.Atoi(data.Price)
		if err != nil {
			log.Println("Harga Bukan String !")
			c.JSON(http.StatusExpectationFailed, gin.H{"message": "Harga Bukan String!", "status": false})
			return
		}

		pcs, err := strconv.Atoi(data.Quantity)
		if err != nil {
			log.Println("Quantity Bukan String !")
			c.JSON(http.StatusExpectationFailed, gin.H{"message": "Quantity Bukan String!", "status": false})
			return
		}

		if data.Diskon == "" {
			data.Diskon = "0"
		}

		disc, err := strconv.Atoi(data.Diskon)
		if err != nil {
			log.Println("Diskon Bukan String !")
			c.JSON(http.StatusExpectationFailed, gin.H{"message": "Diskon Bukan String!", "status": false})

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

func Posting(c *gin.Context) {
	var input model.PurchaseOrder

	if c.GetHeader("content-type") == "application/x-www-form-urlencoded" || c.GetHeader("content-type") == "application/x-www-form-urlencoded; charset=utf-8" {

		if err := c.Bind(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	} else {

		if err := c.BindJSON(&input); err != nil {
			utility.ResponseError(c, "Input Data Tidak Berhasil !")
		}

	}

	ctx := context.Background()
	tx, err := DBConnect.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(ctx)

	log.Println("Data Input :", input)

	jsonData, err := json.Marshal(input)
	if err != nil {
		utility.ResponseError(c, "Gagal mengkonversi data ke JSON!")
		return
	}

	jsonString := string(jsonData)

	log.Println("Json String Input :", jsonString)

	// Generate format PENG+DDMMYYYYHHMMSSMS
	now := time.Now()
	poId := fmt.Sprintf("PENG%02d%02d%d%02d%02d%02d%03d",
		now.Day(), now.Month(), now.Year(),
		now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)

	log.Println("Data Indektik : ", poId)

	query := `
		INSERT INTO pengeluaran (nama, sub_total, pajak, tanggal, total, status, po_id, free_data)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	// Eksekusi query
	_, err = tx.Exec(context.Background(), query, input.NamaSuplier, utility.RupiahToNumber(input.SubTotal), utility.RupiahToNumber(input.Pajak), time.Now().Format("02-01-2006"), utility.RupiahToNumber(input.Total), "DITERIMA", poId, jsonString)
	if err != nil {
		log.Println("Failed to insert data: %v\n", err)
		utility.ResponseError(c, constanta.ErrQuery1)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		utility.ResponseError(c, constanta.ErrCommit)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data added successfully", "status": true})
}
