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

func DeleteItemBuyer(c *gin.Context) {
	var input model.RequestID
	var idcheck int

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

	chek_data := `
		select count(*) from item_buyer where id = $1
	`

	err = tx.QueryRow(ctx, chek_data, input.ID).Scan(idcheck)

	if idcheck == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Data Delete Item successfully but deleted its not from DB", "status": true})
		return
	}

	queryDelete := `
		DELETE FROM item_buyer
		WHERE id = $1;
	`
	_, err = tx.Exec(ctx, queryDelete, input.ID)

	if err != nil {
		tx.Rollback(ctx)
		utility.ResponseError(c, constanta.ErrQuery2)
		return
	}

	if err := tx.Commit(ctx); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Delete Item successfully", "status": true})
}
