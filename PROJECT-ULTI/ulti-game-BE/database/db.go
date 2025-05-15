package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var Database *sql.DB

func InitDatabase() {
	datab, err := sql.Open("sqlite", "tables.sql")

	if err != nil {
		panic("Database could't connect!" + err.Error())
	}

	Database = datab

	err = creatTables()

	if err != nil {
		panic("Database failure!" + err.Error())
	}

	fmt.Println("Tables have been created")
}

func creatTables() error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		state TEXT NOT NULL
	)
	`

	_, err := Database.Exec(createUserTable)

	if err != nil {
		panic("Database could not connect: user" + err.Error())
	}
	return err
}
