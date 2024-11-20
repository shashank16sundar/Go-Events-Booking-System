package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "api.sql")

	if err != nil {
		panic("Couldnt connect to DB")
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateDatabaseTables()
}

func CreateDatabaseTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Couldnt create events table")
	}
}
