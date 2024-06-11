package main

import (
	"backend_project_fismed/database"
	"backend_project_fismed/route"
	"log"

	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-requested-with")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	database.NewConnect()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(CORSMiddleware())

	route.Routes(router)

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("[--->] Running On Port", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
