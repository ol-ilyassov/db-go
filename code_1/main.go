package main

import (
	"log/slog"
	"os"
)

func main() {
	var (
		err error
	)

	// * Подключение к БД.
	db, err := dbconnect()
	if err != nil {
		slog.Error("dbconnect error", "error", err)
		os.Exit(1)
	}
	defer db.Close() // !

	// * Создание таблицы.
	// err = createTableVegetables(db)
	// if err != nil {
	// 	slog.Error("create table error", "error", err)
	// 	os.Exit(1)
	// }

	// * Удаление таблицы.
	// err = dropTableVegetables(db)
	// if err != nil {
	// 	slog.Error("drop table error", "error", err)
	// 	os.Exit(1)
	// }

	// * Получить множество записей.
	// vs, err := getVegetables(db)
	// if err != nil {
	// 	slog.Error("get vegetables error", "error", err)
	// } else {
	// 	slog.Info("get vegetables", "vegetables", vs)
	// }

	// * Добавить запись.
	// vegetable := Vegetable{
	// 	Name:  "tomato",
	// 	Count: 10,
	// 	Price: 800,
	// }
	// err = addVegetable(db, vegetable)
	// if err != nil {
	// 	slog.Error("add vegetable error", "error", err)
	// }

	// * Получить одну запись.
	// v, err := getVegetable(db, "tomato")
	// if err != nil {
	// 	if errors.Is(err, ErrNoRecord) {
	// 		slog.Info("get vegetable", "tomato", ErrNoRecord)
	// 	} else {
	// 		slog.Error("get vegetable error", "error", err)
	// 	}
	// } else {
	// 	slog.Info("get vegetable", "tomato", v)
	// }

	// * Удалить запись.
	// err = removeVegetable(db, "tomato")
	// if err != nil {
	// 	slog.Error("remove vegetable error", "error", err)
	// }

	// * Изменить количество.
	// err = updateCount(db, "tomato", 4)
	// if err != nil {
	// 	slog.Error("update count error", "error", err)
	// }

	// * Изменить цену.
	// err = setPrice(db, "tomato", 900)
	// if err != nil {
	// 	slog.Error("set price error", "error", err)
	// }

	slog.Info("program finished")
}
