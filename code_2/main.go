package main

import (
	"context"
	"dbgo/code_2/db"
	"dbgo/code_2/repo"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var (
		err error
		ctx = context.Background()
	)

	conn, err := pgxpool.New(ctx, db.DSN) // pool of connections.
	if err != nil {
		slog.Error("pgxpool.New", "error", err)
		os.Exit(1)
	}
	defer conn.Close() // ?

	repo := repo.New(conn)

	book, err := repo.Get(ctx, 1)
	if err != nil {
		slog.Error("repo.Get", "error", err)
		os.Exit(1)
	} else {
		slog.Info("repo.Get", "book", book)
	}
}
