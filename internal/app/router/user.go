package router

import (
	"gin-design/internal/app/controller"

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
	user := router.Group("/user")
	user.GET("", r.ctrl.GetUser)
}
