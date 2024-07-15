package database

import (
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"backend_project_fismed/service/pemasukan"
	"backend_project_fismed/service/piutang"
	"backend_project_fismed/service/preOrder"
	"backend_project_fismed/service/proformaInvoice"
	"backend_project_fismed/service/salesOrder"
	"backend_project_fismed/service/stockBarang"
	"backend_project_fismed/utility"
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnect() *pgxpool.Pool {
	//	DB LOCALHOST ADIT
	//databaseUrl := "postgres://postgres:admin123@localhost:5432/fismed"

	//  DB FISMED DEV
	//databaseUrl := "postgres://fismed-user:fismed-db-12345@62.146.233.39:5445/fismed"

	//  DB FISMED PRODUCTION
	databaseUrl := "postgres://fismed-user:fismed-db-12345@209.182.237.155:5445/fismed"

	//  BUAT DOCKER
	//databaseUrl := "postgres://fismed-user:boyang123@fismed-db:5432/fismed"

	//db, err := pgxpool.Connect(context.Background(), databaseUrl)
	//if err != nil {
	//	log.Fatalf("Unable to connect to database: %v\n", err)
	//}

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
		os.Exit(1)
	}

	config.MaxConns = 100
	config.ConnConfig.ConnectTimeout = 2 * time.Second
	config.HealthCheckPeriod = 2 * time.Second

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	log.Println("[--->] Success Created DB Connection...!")

	authentikasi.InitiateDB(db)
	utility.InitiateDB(db)
	customerProfilling.InitiateDB(db)
	proformaInvoice.InitiateDB(db)
	stockBarang.InitiateDB(db)
	preOrder.InitiateDB(db)
	salesOrder.InitiateDB(db)
	pemasukan.InitiateDB(db)
	piutang.InitiateDB(db)

	return db
}
