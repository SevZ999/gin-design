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
	NewAccessRouter,
)

func NewRouters(
	user *UserRouter,
	access *AccessRouter,
) []Router {
	return []Router{
		user,
		access,
	}
}
