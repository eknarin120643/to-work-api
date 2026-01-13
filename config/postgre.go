package config

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type PostgreConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgreConfig(host, port, user, password, dbName string) (*PostgreConfig, error) {
	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		return nil, errors.New("env variables are not set")
	}
	return &PostgreConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}

func (p *PostgreConfig) InitPostgre() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Warning: .env file not found, using system environment variables", "error", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DBName)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	slog.Info("Successfully connected to PostgreSQL")
	return pool, nil
}
