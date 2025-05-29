// middleware/logging.go
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 修改函数签名接收*zap.Logger
func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetString("request_id")

		c.Next()

		status := c.Writer.Status()
		latency := time.Since(start)
		err := c.Errors.Last()

		fields := []zap.Field{
			zap.String("request_id", requestID),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.ClientIP()),
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
		}

		if status >= 500 {
			logger.Error("Internal server error", fields...)
		} else if status >= 400 {
			logger.Warn("Client error", fields...)
		} else {
			logger.Info("Request completed", fields...)
		}
	}
}
