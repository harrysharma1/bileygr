package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DevDB *sql.DB

func InitDevDB() {
	var err error
	uri := "postgresql://harrysharma@localhost/bileygr"
	DevDB, err = sql.Open("pgx", uri)
	if err != nil {
		log.Fatalf("error connecting to database: %s\n", err)
	}

	if err := DevDB.Ping(); err != nil {
		log.Fatalf("error pinging database: %s\n", err)
	}
}
