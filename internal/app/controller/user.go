package controller

import (
	"gin-design/internal/app/dto"
	"gin-design/internal/app/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		srv: userService,
	}
}

func (c *UserController) GetUser(ctx *gin.Context) {

	// resp, err := c.srv.GetUser(GetCtx(ctx), dto.GetUserReq{Id: 200})
	// if err != nil {
	// 	ctx.JSON(200, dto.Error(
	// 		1,
	// 		err.Error(),
	// 	))
	// }
	// ctx.JSON(200, dto.Success(resp))
	ctx.JSON(200, dto.Success("this is gray"))
}
