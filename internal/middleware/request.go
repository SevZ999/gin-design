package middleware

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 尝试从 Header 获取现有 ID
		requestID := c.GetHeader("request-id")

		// 2. 如果不存在则生成新 ID
		if requestID == "" {
			requestID = generateRequestId()
			c.Writer.Header().Set("request-id", requestID)
		}

		// 3. 存储到上下文供后续使用
		c.Set("request-id", requestID)
		c.Next()
	}
}

func generateRequestId() string {
	// 使用 UUIDv4 保证随机性
	id, err := uuid.NewRandom()
	if err != nil {
		// 回退方案：时间戳 + 随机数
		return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
	}
	return id.String()
}
