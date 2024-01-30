package main

import (
	"log"

	"github.com/felipefbs/ick-app/database"
	"github.com/felipefbs/ick-app/server"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	server := server.Init(db)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
