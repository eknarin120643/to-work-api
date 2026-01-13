package main

import (
	"to-work-api/cmd/rest/handler"
	"to-work-api/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(mode config.Mode, h *handler.Handler) *gin.Engine {
	if mode == config.ModeRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

	// Trusted Proxies
	r.SetTrustedProxies(nil)
	v1 := r.Group("/api/v1")
	v1.GET("/ping", h.Ping)

	return r
}
