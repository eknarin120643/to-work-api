package main

import (
	"log/slog"

	"to-work-api/cmd/rest/handler"
	"to-work-api/config"
)

func main() {
	// load config
	slog.Info("Loading config...")
	config, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error config", err)
	}
	slog.Info("Config loaded successfully")

	// init db
	slog.Info("Initializing database...")
	db, err := config.PostgreConfig.InitPostgre()
	if err != nil {
		slog.Error("Failed to initialize database", "error db", err)
	}
	slog.Info("Database initialized successfully")

	// init redis
	slog.Info("Initializing redis...")
	redis, err := config.RedisConfig.InitRedis()
	if err != nil {
		slog.Error("Failed to initialize redis", "error redis", err)
	}
	slog.Info("Redis initialized successfully")

	handler := handler.NewHandler(db, redis)
	startServer(handler)
}
