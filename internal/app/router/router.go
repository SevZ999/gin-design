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
	NewShopRouter,
)

func NewRouters(
	user *UserRouter,
	shop *ShopRouter,
) []Router {
	return []Router{
		user,
		shop,
	}
}
