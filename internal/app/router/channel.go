package router

import (
	"loan-admin/internal/app/controller"

	"github.com/gin-gonic/gin"
)

type ChannelRouter struct {
	ctrl *controller.ChannelController
}

func NewChannelRouter(ctrl *controller.ChannelController) *ChannelRouter {
	return &ChannelRouter{
		ctrl: ctrl,
	}
}

func (r *ChannelRouter) SetRoute(router *gin.RouterGroup) {
	router.GET("/channel", r.ctrl.GetChannel)
}
