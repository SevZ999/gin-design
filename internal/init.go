package internal

import (
	"gin-design/internal/app/router"
	"gin-design/internal/config"
	"gin-design/internal/middleware"
	"gin-design/internal/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

var isReady = true

func NewEngine(cfg *config.Config, logger *logger.ZapLogger, routers []router.Router) *gin.Engine {
	gin.SetMode(cfg.Env)

	r := gin.Default()

	r.GET("/healthz", kubernetesHealthz())
	r.GET("/readyz", kubernetesReadyz())
	r.GET("/shutdown", kubernetesShutdown())

	//全局中间件
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

func kubernetesHealthz() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func kubernetesReadyz() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isReady {
			c.Status(http.StatusOK) // 就绪时返回200
		} else {
			c.Status(http.StatusServiceUnavailable) // 终止前返回503，触发就绪探针失效
		}
	}
}

func kubernetesShutdown() gin.HandlerFunc {
	return func(c *gin.Context) {
		isReady = false // 标记为未就绪，后续/readyz返回503
		c.Status(http.StatusOK)
	}
}
