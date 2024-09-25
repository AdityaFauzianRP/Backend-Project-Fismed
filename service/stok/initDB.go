package stok

import "github.com/jackc/pgx/v4/pgxpool"

var DBConnect *pgxpool.Pool

func InitiateDB(db *pgxpool.Pool) {

	DBConnect = db

}
