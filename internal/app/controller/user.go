package controller

import (
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/service"

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

func (ctrl *UserController) Check(c *gin.Context) {
	requestId := c.Request.Header.Get("X-Request-ID")
	authorization := c.Request.Header.Get("Authorization")

	c.JSON(200, dto.Success(
		map[string]interface{}{
			"request_id":    requestId,
			"authorization": authorization,
		},
	))
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(400, dto.Error(400, "invalid id"))
	// 	return
	// }
	resp, err := ctrl.srv.GetUser(1)
	if err != nil {
		c.JSON(500, dto.Error(500, err.Error()))
		return
	}

	c.JSON(200, dto.Success(resp))
}

func (ctrl *UserController) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.Error(400, err.Error()))
		return
	}
	resp, err := ctrl.srv.Login(req)
	if err != nil {
		c.JSON(500, dto.Error(500, err.Error()))
		return
	}
	c.JSON(200, dto.Success(resp))
}
