package main

import (
	"database/sql" // стандарт: интерфейс подключения к базе данных.
	"errors"
	"fmt"

	_ "github.com/lib/pq" // драйвер подключения к базе данных.
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dbuser"
	password = "dbpassword"
	dbname   = "dbgo"
	sslmode  = "disable"
)

// Подключение к базе данных.
func dbconnect() (*sql.DB, error) {
	// dsn = data source name.
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	// db - набор множества подключений к базе данных.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

// * Data Definition Language:

// Запрос на создание таблицы.
func createTableVegetables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS vegetables(
			name text, 
			count int4, 
			price float8
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// Запрос на удаление таблицы.
func dropTableVegetables(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS vegetables;`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// * Data Manipulation Language:

// Запрос на получение всех записей.
func getVegetables(db *sql.DB) ([]*Vegetable, error) {
	query := `
		SELECT name, count, price
		FROM vegetables
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // !

	var vegetables []*Vegetable

	if rows.Next() {
		var v Vegetable // ?

		err = rows.Scan(&v.Name, &v.Count, &v.Price)
		if err != nil {
			return nil, err
		}

		vegetables = append(vegetables, &v)
	}
	if err = rows.Err(); err != nil { // !
		return nil, err
	}

	return vegetables, nil
}

// Запрос на добавление записи.
func addVegetable(db *sql.DB, v Vegetable) error {
	query := `
		INSERT INTO vegetables(name, count, price)
		VALUES ($1, $2, $3)
	`
	args := []any{v.Name, v.Count, v.Price}

	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Запрос на получение одной записи по названию (name).
func getVegetable(db *sql.DB, name string) (*Vegetable, error) {
	query := `
		SELECT name, count, price
		FROM vegetables
		WHERE name = $1
	`

	var v Vegetable

	row := db.QueryRow(query, name)

	err := row.Scan(&v.Name, &v.Count, &v.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return &v, nil
}

// Запрос на удаление записи.
func removeVegetable(db *sql.DB, name string) error {
	query := `
		DELETE FROM vegetables
		WHERE name = $1
	`

	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

// Запрос на изменение значения количества записи.
func updateCount(db *sql.DB, name string, count int) error {
	query := `
		UPDATE vegetables
		SET count = count + $1
		WHERE name = $2
	`

	_, err := db.Exec(query, count, name)
	if err != nil {
		return err
	}

	return nil
}

// Запрос на изменение значения цены записи.
func setPrice(db *sql.DB, name string, price float64) error {
	query := `
		UPDATE vegetables
		SET price = $1
		WHERE name = $2
	`

	_, err := db.Exec(query, price, name)
	if err != nil {
		return err
	}

	return nil
}
