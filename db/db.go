package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DevDB *sql.DB

type Database struct {
	DB *sql.DB
}

func NewDatabase(connectionString string) (*Database, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil

}

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
