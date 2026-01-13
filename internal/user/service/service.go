package service

import (
	"to-work-api/internal/sqlc"
	"to-work-api/internal/user/service/query"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	// TODO : implement commands
}

type Queries struct {
	GetUser query.GetUserHandler
}

func NewService(db *pgxpool.Pool, rc *redis.Client) *Service {
	queries := sqlc.New(db)
	return &Service{
		Commands: Commands{},
		Queries: Queries{
			GetUser: *query.NewGetUserHandler(queries),
		},
	}
}
