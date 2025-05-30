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
	NewChannelRouter,
	NewShopRouter,
)

func NewRouters(
	user *UserRouter,
	access *AccessRouter,
	channel *ChannelRouter,
	shop *ShopRouter,
) []Router {
	return []Router{
		user,
		access,
		channel,
		shop,
	}
}
