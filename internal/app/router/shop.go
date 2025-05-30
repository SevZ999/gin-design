package router

import (
	"loan-admin/internal/app/controller"

	"github.com/gin-gonic/gin"
)

type ShopRouter struct {
	ctrl *controller.ShopController
}

func NewShopRouter(ctrl *controller.ShopController) *ShopRouter {
	return &ShopRouter{
		ctrl: ctrl,
	}
}

func (r *ShopRouter) SetRoute(router *gin.RouterGroup) {
	api := router.Group("/shop")
	api.GET("", r.ctrl.GetShops)
}
