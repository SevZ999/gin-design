package controller

import (
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	srv *service.ChannelService
}

func NewChannelController(srv *service.ChannelService) *ChannelController {
	return &ChannelController{srv: srv}
}

func (ctrl *ChannelController) GetChannel(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Success(dto.GetChannelResp{
		Channels: []dto.Channel{
			{
				Id:   1,
				Name: "渠道1",
			},
			{
				Id:   2,
				Name: "渠道2",
			},
			{
				Id:   3,
				Name: "渠道3",
			},
		},
	}))
}
