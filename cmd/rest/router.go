package main

import (
	"to-work-api/cmd/rest/handler"
	"to-work-api/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(mode config.Mode, h *handler.Handler) *gin.Engine {
	if mode == config.ModeRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/ping", h.Ping)

	return r
}
