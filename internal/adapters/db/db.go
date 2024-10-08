package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
)

type DBAdapter struct {
	conn    *pgxpool.Pool
	queries *queries.Queries
}

func NewDB(DBUrl string) *DBAdapter {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, DBUrl)
	if err != nil {
		log.Fatal("Error connecting to db:%w", err)
	}
	// defer conn.Close(ctx)

	queries := queries.New(conn)

	return &DBAdapter{conn: conn, queries: queries}

}


