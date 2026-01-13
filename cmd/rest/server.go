package main

import (
	"log/slog"

	"to-work-api/cmd/rest/handler"
	"to-work-api/config"
)

func startServer(mode config.Mode, h *handler.Handler) {
	r := NewRouter(mode, h)
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to run server", "error", err)
	}
}
