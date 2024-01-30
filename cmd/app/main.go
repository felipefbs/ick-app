package main

import (
	"context"
	"log"

	"github.com/felipefbs/ick-app/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig("user=admin password=12345 host=localhost port=5432 dbname=icks sslmode=disable pool_max_conns=10")
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgxpool.New(ctx, config.ConnString())
	if err != nil {
		log.Fatal(err)
	}

	server := server.Init(db)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
