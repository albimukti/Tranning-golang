package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(connString string) {
	var err error
	DB, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func CloseDB() {
	DB.Close()
}
