package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Router interface {
	SetRoute(router *gin.RouterGroup)
}

var RouterProviderSet = wire.NewSet(
	NewUserRouter,
)

func NewRouters(
	user *UserRouter,
) []Router {
	return []Router{
		user,
	}
}
