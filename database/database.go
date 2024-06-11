package database

import (
	"backend_project_fismed/service"
	"backend_project_fismed/service/authentikasi"
	"log"

	"os"

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnect() *pgxpool.Pool {
	databaseUrl := "postgres://postgres:admin123@localhost:5432/fismed"
	//databaseUrl := "postgres://postgres:boyang123@morodb.cmwu6s1vldt3.ap-southeast-1.rds.amazonaws.com:5432/morowali"

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
		os.Exit(1)
	}

	config.MaxConns = 10

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	log.Println("[--->] Success Created DB Connection...!")

	authentikasi.InitiateDB(db)
	service.InitiateDB(db)

	return db
}
