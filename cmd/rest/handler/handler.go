package handler

import (
	"net/http"

	userservice "to-work-api/internal/user/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	User *userservice.Service
}

func NewHandler(db *pgxpool.Pool, rc *redis.Client) *Handler {
	return &Handler{
		User: userservice.NewService(db, rc),
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
