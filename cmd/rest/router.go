package main

import (
	"to-work-api/cmd/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/ping", h.Ping)

	return r
}
