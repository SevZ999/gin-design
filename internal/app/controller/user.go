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

func (ctrl *UserController) GetUser(c *gin.Context) {
	var req dto.GetUserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, dto.Error(400, err.Error()))
		return
	}

	resp, err := ctrl.srv.GetUser(req.Id)
	if err != nil {
		c.JSON(500, dto.Error(500, err.Error()))
		return
	}

	c.JSON(200, dto.Success(resp))
}
