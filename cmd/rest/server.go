package main

import (
	"log/slog"

	"to-work-api/cmd/rest/handler"
)

func startServer(h *handler.Handler) {
	r := NewRouter(h)
	slog.Info("REST Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to run server", "error", err)
	}
}
