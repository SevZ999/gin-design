package controller

import (
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/service"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
	srv *service.ShopService
}

func NewShopController(userService *service.ShopService) *ShopController {
	return &ShopController{
		srv: userService,
	}
}

func (ctrl *ShopController) GetShops(c *gin.Context) {
	resp, err := ctrl.srv.GetShop()
	if err != nil {
		c.JSON(500, dto.Error(500, err.Error()))
		return
	}

	c.JSON(200, dto.Success(resp))
}
