package router

import (
	"loan-admin/internal/app/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	ctrl *controller.UserController
}

func NewUserRouter(ctrl *controller.UserController) *UserRouter {
	return &UserRouter{
		ctrl: ctrl,
	}
}

func (r *UserRouter) SetRoute(router *gin.RouterGroup) {
	api := router.Group("/user")

	api.POST("/login", r.ctrl.Login)

	api.GET("", r.ctrl.GetUser)

	api.GET("check", r.ctrl.Check)

}
