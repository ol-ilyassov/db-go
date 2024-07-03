package main

import (
	"log/slog"
	"os"
)

const (
	DSN = "postgres://dbuser:dbpassword@localhost:5432/dbgo?sslmode=disable"
)

func main() {
	err := MigrateUp(DSN)
	if err != nil {
		slog.Error("MigrateUp", "error", err)
		os.Exit(1)
	}
}
