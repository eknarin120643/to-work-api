package config

import (
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Mode string

func (m Mode) String() string {
	return string(m)
}

func newModeFromString(mode string) Mode {
	switch mode {
	case "DEVELOP":
		return ModeDevelop
	case "RELEASE":
		return ModeRelease
	default:
		return ModeDevelop
	}
}

const (
	ModeDevelop Mode = "DEVELOP"
	ModeRelease Mode = "RELEASE"
)

type Config struct {
	Mode          Mode
	PostgreConfig *PostgreConfig
	RedisConfig   *RedisConfig
}

func Load() (*Config, error) {
	// load .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	mode := os.Getenv("MODE")

	// init redis config
	postgreConfig, err := NewPostgreConfig(os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	if err != nil {
		return nil, err
	}

	redisConfig, err := NewRedisConfig(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		return nil, err
	}

	config := &Config{
		Mode:          newModeFromString(mode),
		PostgreConfig: postgreConfig,
		RedisConfig:   redisConfig,
	}

	return config, nil
}

func (c *Config) InitPostgre() (*pgxpool.Pool, error) {
	return c.PostgreConfig.InitPostgre()
}

func (c *Config) InitRedis() (*redis.Client, error) {
	return c.RedisConfig.InitRedis()
}
