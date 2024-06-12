package route

import (
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/login", authentikasi.Login)
	router.POST("/token-validate", authentikasi.TokenValidate)

	//	Customer Profilling API

	router.POST("/customer-profilling/add", customerProfilling.Add)
	router.POST("/customer-profilling/get-tax-code", customerProfilling.GetTaxCode)

}
