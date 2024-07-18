package main

import (
	"backend_project_fismed/database"
	"backend_project_fismed/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS Middleware hit - Before setting headers")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*, x-requested-with")

		log.Println("CORS Middleware hit - Before c.Next()")
		//if c.Request.Method == "OPTIONS" {
		//	log.Println("CORS Middleware hit - OPTIONS request")
		//	c.AbortWithStatus(http.StatusNoContent)
		//	return
		//}

		// Set timeout for handling the request
		//ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		//defer cancel()

		//c.Request = c.Request.WithContext(ctx)

		//c.Next()

		log.Println("CORS Middleware hit - After c.Next()")
	}
}

func main() {
	database.NewConnect()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	route.Routes(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	log.Printf("[--->] Running On Port %s", server.Addr)

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
