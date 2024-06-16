package database

import (
	"backend_project_fismed/service"
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"backend_project_fismed/service/proformaInvoice"
	"backend_project_fismed/service/stockBarang"
	"log"

	"os"

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnect() *pgxpool.Pool {
	databaseUrl := "postgres://postgres:admin123@localhost:5432/fismed"
	//buat docker
	//databaseUrl := "postgres://fismed-user:boyang123@fismed-db:5432/fismed"

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
		os.Exit(1)
	}

	config.MaxConns = 2

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	log.Println("[--->] Success Created DB Connection...!")

	authentikasi.InitiateDB(db)
	service.InitiateDB(db)
	customerProfilling.InitiateDB(db)
	proformaInvoice.InitiateDB(db)
	stockBarang.InitiateDB(db)

	return db
}
