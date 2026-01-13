package config

import (
	"context"
	"errors"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisConfig(addr, password string) (*RedisConfig, error) {
	if addr == "" || password == "" {
		return nil, errors.New("env variables are not set")
	}
	return &RedisConfig{
		Addr:     addr,
		Password: password,
		DB:       0,
	}, nil
}

func (r *RedisConfig) InitRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB, // use from input
	})

	// Test the connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	slog.Info("Successfully connected to Redis")
	return rdb, nil
}
