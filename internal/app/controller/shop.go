package controller

import (
	"gin-design/internal/app/service"

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

func (ctrl *ShopController) GetShop(c *gin.Context) {
	c.JSON(200, "ok")
}
