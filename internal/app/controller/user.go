package controller

import (
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} dto.GetUserResp
// @Failure 400 {string} string "无效请求"
// @Failure 500 {string} string "服务器错误"
// @Router /api/user/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(400, dto.Error(400, "invalid id"))
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, dto.Error(400, "invalid id"))
		return
	}

	resp, err := ctrl.srv.GetUser(userId)
	if err != nil {
		c.JSON(500, dto.Error(500, err.Error()))
		return
	}

	c.JSON(200, dto.Success(resp))
}

func (ctrl *UserController) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Sugar().Error(err)
		c.JSON(400, dto.Error(400, "invalid request"))
		return
	}
	resp, err := ctrl.srv.Login(req)
	if err != nil {
		c.JSON(500, dto.Error(500, err.Error()))
		return
	}
	c.JSON(200, dto.Success(resp))
}
