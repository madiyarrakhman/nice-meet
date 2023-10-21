package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Initialize() (*sql.DB, error) {
	return sql.Open("sqlite3", "./database.sqlite")
}
