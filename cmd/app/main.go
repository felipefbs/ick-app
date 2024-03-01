package main

import (
	"log/slog"

	"github.com/felipefbs/ick-app/internal/server"
)

func main() {

	server := server.Init()

	if err := server.ListenAndServe(); err != nil {
		slog.Error("failed to init http server", "error", err)
	}
}
