package utils

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateRequestId() string {
	// 使用 UUIDv4 保证随机性
	id, err := uuid.NewRandom()
	if err != nil {
		// 回退方案：时间戳 + 随机数
		return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
	}
	return id.String()
}

func GetRequestId(ctx context.Context) string {
	requestId, ok := ctx.Value(gin.ContextRequestKey).(string)
	if ok {
		return requestId
	} else {
		return ""
	}
}
