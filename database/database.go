package database

import (
	"backend_project_fismed/service/authentikasi"
	"backend_project_fismed/service/customerProfilling"
	"backend_project_fismed/service/pemasukan"
	"backend_project_fismed/service/pengeluaran"
	"backend_project_fismed/service/piutang"
	"backend_project_fismed/service/preOrder"
	price_list "backend_project_fismed/service/price-list"
	"backend_project_fismed/service/proformaInvoice"
	"backend_project_fismed/service/salesOrder"
	"backend_project_fismed/service/stockBarang"
	"backend_project_fismed/service/stok"
	"backend_project_fismed/utility"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var db *pgxpool.Pool

func NewConnect() (*pgxpool.Pool, error) {
	//  DB FISMED PRODUCTION
	databaseUrl := "postgres://fismed-user:fismed-db-12345@209.182.237.155:5445/fismed"

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
		return nil, err
	}

	config.MaxConns = 1000
	config.ConnConfig.ConnectTimeout = 20 * time.Second
	config.HealthCheckPeriod = 1 * time.Minute

	db, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
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
	pengeluaran.InitiateDB(db)
	stok.InitiateDB(db)
	price_list.InitiateDB(db)

	return db, nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}
