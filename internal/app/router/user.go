package router

import (
	"gin-design/internal/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	ctrl *controller.UserController
}

func NewUserRouter(ctrl *controller.UserController) *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) Route(r *gin.RouterGroup) {
	r.Use()

	user := r.Group("/user")

	user.GET("", u.ctrl.GetUser)
}
