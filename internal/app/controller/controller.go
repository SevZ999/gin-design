package controller

import (
	"context"
	"gin-design/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// 定义自定义键类型（避免冲突）

var ControllerProviderSet = wire.NewSet(
	NewUserController,
	NewShopController,
)

func GetCtx(c *gin.Context) context.Context {
	requestID := c.GetString("request-id")
	if requestID == "" {
		return context.WithValue(c.Request.Context(), gin.ContextRequestKey, utils.GenerateRequestId()) // 返回原始请求上下文（不携带requestID）
	}
	// 基于原始请求上下文创建携带requestID的子上下文
	return context.WithValue(c.Request.Context(), gin.ContextRequestKey, requestID)
}
