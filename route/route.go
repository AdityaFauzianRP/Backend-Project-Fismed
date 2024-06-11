package route

import (
	"backend_project_fismed/service/authentikasi"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/login", authentikasi.Login)
	router.POST("/token-validate", authentikasi.TokenValidate)

}
