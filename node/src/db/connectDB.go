package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (database *sql.DB) {
	database, _ = sql.Open("sqlite3", "./db/database.db")
	return
}
