package database

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dbPool *sql.DB
var dbMutex = &sync.Mutex{}

func Get() (*sql.DB, error) {
	if dbPool == nil {
		dbMutex.Lock()
		defer dbMutex.Unlock()

		if dbPool == nil {
			slog.Info("initializing database connection")

			url := os.Getenv("DB_URL") + "?authToken=" + os.Getenv("DB_TOKEN")

			db, err := sql.Open("libsql", url)
			if err != nil {
				slog.Error("failed to load database", "error", err)

				return nil, err
			}

			dbPool = db
		}
	}

	return dbPool, nil
}
