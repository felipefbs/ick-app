package database

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func InitDatabase() (*sql.DB, error) {
	return sql.Open("sqlite", "local.db")
}
