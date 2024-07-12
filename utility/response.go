package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, responseMessage string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error":  responseMessage,
		"status": false,
	})

	return
}
