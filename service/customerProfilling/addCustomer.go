package customerProfilling

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func Add(c *gin.Context) {
	type Request struct {
		Name               string `json:"name"`
		AddressCompany     string `json:"address_company"`
		NPWPAddress        string `json:"npwp_address"`
		NPWP               string `json:"npwp"`
		IpakNumber         string `json:"ipak_number"`
		FactureAddress     string `json:"facture_address"`
		CityFacture        string `json:"city_facture"`
		ZipCodeFacture     string `json:"zip_code_facture"`
		NumberPhoneFacture string `json:"number_phone_facture"`
		EmailFacture       string `json:"email_facture"`
		FaxFacture         string `json:"fax_facture"`
		PicFacture         string `json:"pic_facture"`
		ItemAddress        string `json:"item_address"`
		CityItem           string `json:"city_item"`
		ZipCodeItem        string `json:"zip_code_item"`
		NumberPhoneItem    string `json:"number_phone_item"`
		EmailItem          string `json:"email_item"`
		FaxItem            string `json:"fax_item"`
		PicItem            string `json:"pic_item"`
		ContactPerson      string `json:"contact_person"`
		TaxCodeID          string `json:"tax_code_id"`
		Top                string `json:"top"`
	}

	var input Request

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

	query := `INSERT INTO customer (
		name, address_company, npwp_address, npwp, ipak_number, 
		facture_address, city_facture, zip_code_facture, number_phone_facture, email_facture, fax_facture, pic_facture,
		item_address, city_item, zip_code_item, number_phone_item, email_item, fax_item, pic_item,
		contact_person, tax_code_id, top,
		created_at, created_by, updated_at, updated_by
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, NOW(), 'admin', NOW(), 'admin'
	)`

	_, err = tx.Exec(ctx, query,
		input.Name, input.AddressCompany, input.NPWPAddress, input.NPWP, input.IpakNumber,
		input.FactureAddress, input.CityFacture, input.ZipCodeFacture, input.NumberPhoneFacture, input.EmailFacture, input.FaxFacture, input.PicFacture,
		input.ItemAddress, input.CityItem, input.ZipCodeItem, input.NumberPhoneItem, input.EmailItem, input.FaxItem, input.PicItem,
		input.ContactPerson, input.TaxCodeID, input.Top)

	if err != nil {
		tx.Rollback(ctx)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "status": false})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully", "status": true})

}
