package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitDB() {
	uri := "postgresql://harrysharma@localhost:5432/bileygr"
	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		log.Fatalf("error connecting to db: %e", err)
	}
	defer conn.Close(context.Background())
}
