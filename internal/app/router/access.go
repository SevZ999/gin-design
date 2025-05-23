package router

import (
	"loan-admin/internal/app/controller"

	"github.com/gin-gonic/gin"
)

type AccessRouter struct {
	ctrl *controller.AccessController
}

func NewAccessRouter(ctrl *controller.AccessController) *AccessRouter {
	return &AccessRouter{
		ctrl: ctrl,
	}
}

func (r *AccessRouter) SetRoute(router *gin.RouterGroup) {

	api := router.Group("/access")

	api.GET("", r.ctrl.GetAccess)
}
