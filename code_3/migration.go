package main

import (
	"embed"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// Run executes migration up scripts on SQLite3.
func MigrateUp(dbstring string) (err error) {
	driver := "postgres"
	dir := "migrations"
	goose.SetTableName("version")
	goose.SetBaseFS(embedMigrations)

	db, err := goose.OpenDBWithDriver(driver, dbstring)
	if err != nil {
		return err
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			err = errClose
			return
		}
	}()

	if err = goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}
