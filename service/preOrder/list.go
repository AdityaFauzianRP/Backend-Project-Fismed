package preOrder

import (
	"backend_project_fismed/request"
	"github.com/gin-gonic/gin"
	"log"
)

func ListPO(c *gin.Context) {
	requestData, err := request.ParseRequest(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Simpan data request ke dalam variabel global
	request.GlobalRequestData = requestData

	// Konversi bagian tertentu dari body menjadi tipe yang sesuai
	reqData, err := request.InterfaceToMapOrSliceOfMap(request.GlobalRequestData.Body["json"])
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println("Request Data :", reqData)
}
