package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	*pgxpool.Pool
}

func NewDatabase() *Database {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Database{
		dbpool,
	}
}
