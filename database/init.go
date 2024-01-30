package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDatabase() (*sql.DB, error) {
	return sql.Open("sqlite", "local.db")
}
