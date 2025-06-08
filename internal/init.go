package internal

import (
	"gin-design/internal/app/router"
	"gin-design/internal/config"
	"gin-design/internal/middleware"
	"gin-design/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewEngine(cfg *config.Config, logger *logger.ZapLogger, routers []router.Router) *gin.Engine {
	gin.SetMode(cfg.Env)

	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/readyz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(
		middleware.CORS(),
		middleware.RequestId(),
		middleware.LoggingMiddleware(logger),
	)

	api := r.Group("api")

	setRoute(api, routers)

	return r
}

func setRoute(api *gin.RouterGroup, routers []router.Router) {
	for _, r := range routers {
		r.SetRoute(api)
	}
}
