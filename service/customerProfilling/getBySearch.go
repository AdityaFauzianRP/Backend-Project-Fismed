package customerProfilling

import (
	"backend_project_fismed/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
)

func GetBySearch(c *gin.Context) {

	var input model.Customer

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

	log.Println("[--->]", "Data Input :", input)

	query := `
		SELECT 
			id,
			COALESCE("name", '') AS "name",
			COALESCE(address_company, '') AS address_company,
			COALESCE(npwp_address, '') AS npwp_address,
			COALESCE(npwp, '') AS npwp,
			COALESCE(ipak_number, '') AS ipak_number,
			COALESCE(facture_address, '') AS facture_address,
			COALESCE(city_facture, '') AS city_facture,
			COALESCE(zip_code_facture, '') AS zip_code_facture,
			COALESCE(number_phone_facture, '') AS number_phone_facture,
			COALESCE(email_facture, '') AS email_facture,
			COALESCE(fax_facture, '') AS fax_facture,
			COALESCE(pic_facture, '') AS pic_facture,
			COALESCE(item_address, '') AS item_address,
			COALESCE(city_item, '') AS city_item,
			COALESCE(zip_code_item, '') AS zip_code_item,
			COALESCE(number_phone_item, '') AS number_phone_item,
			COALESCE(email_item, '') AS email_item,
			COALESCE(fax_item, '') AS fax_item,
			COALESCE(pic_item, '') AS pic_item,
			COALESCE(contact_person, '') AS contact_person,
			COALESCE(CAST(tax_code_id AS TEXT), '0') AS tax_code_id,  -- Assuming tax_code_id is an integer. Adjust the default value as needed.
			COALESCE(top, '') AS top
		FROM 
			public.customer
		where name = $1;
	`

	rows, err := tx.Query(ctx, query, input.Name)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute tax code query", "status": false})
		return
	}
	defer rows.Close()

	var response []model.Customer

	for rows.Next() {
		var tc model.Customer
		if err := rows.Scan(
			&tc.ID,
			&tc.Name,
			&tc.AddressCompany,
			&tc.NPWPAddress,
			&tc.NPWP,
			&tc.IpakNumber,
			&tc.FactureAddress,
			&tc.CityFacture,
			&tc.ZipCodeFacture,
			&tc.NumberPhoneFacture,
			&tc.EmailFacture,
			&tc.FaxFacture,
			&tc.PicFacture,
			&tc.ItemAddress,
			&tc.CityItem,
			&tc.ZipCodeItem,
			&tc.NumberPhoneItem,
			&tc.EmailItem,
			&tc.FaxItem,
			&tc.PicItem,
			&tc.ContactPerson,
			&tc.TaxCodeID,
			&tc.Top,
		); err != nil {
			log.Println("[--->]", "Error Scan Data :", error(err))
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Data", "status": false})
			return
		}
		response = append(response, tc)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over tax code rows", "status": false})
		return
	}

	var jumlahPI int

	countPI := `
		SELECT COUNT(*) AS total_invoices
		FROM performance_invoice
		WHERE customer_id = $1;
	`

	err = tx.QueryRow(ctx, countPI, response[0].ID).Scan(&jumlahPI)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute tax code query", "status": false})
		return
	}

	log.Println("[--->]", "Nama Customer :", response[0].Name)

	QueryHistory := `
		select 
		    a."name" , 
		    a.quantity  
		from 
		    order_items a, performance_invoice b 
		where  
		    b.customer_id = $1 and 
		    b.id = a.pi_id;
	`

	rows, err = tx.Query(ctx, QueryHistory, response[0].ID)
	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute tax code query", "status": false})
		return
	}
	defer rows.Close()

	var responseRiwayat []model.StockBarang

	for rows.Next() {
		var tc model.StockBarang
		if err := rows.Scan(
			&tc.Name,
			&tc.Total,
		); err != nil {
			log.Println("[--->]", "Error Scan Data :", error(err))
			tx.Rollback(ctx)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Data", "status": false})
			return
		}
		responseRiwayat = append(responseRiwayat, tc)
	}

	log.Println("[--->]", "Data Response Riwayat :", responseRiwayat)

	itemQuantities := make(map[string]int)

	// Iterate over the order items and accumulate the quantities for each item name
	for _, item := range responseRiwayat {
		total, err := strconv.Atoi(item.Total)
		if err != nil {
			log.Println("[--->]", "Count Data Error :", error(err))
		}
		itemQuantities[item.Name] += total
	}

	// Print the total quantity for each item name
	for itemName, quantity := range itemQuantities {
		fmt.Printf("Item: %s, Total Quantity: %d\n", itemName, quantity)
	}

	var itemTotals []model.StockBarang

	for itemName, quantity := range itemQuantities {
		itemTotals = append(itemTotals, model.StockBarang{Name: itemName, Total: strconv.Itoa(quantity)})
	}

	if len(response) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":       "Data Ditemukan !",
			"jumlah_pi":     jumlahPI,
			"customer":      response,
			"riwayat_order": itemTotals,
			"status":        true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Tidak Ditemukan !", "Data": []model.Customer{}, "status": true})
	}

}
