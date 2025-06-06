package service

import "gin-design/internal/pkg/logger"

type ShopRepo interface {
}

type ShopService struct {
	repo ShopRepo
	log  logger.Logger // 新增日志字段
}

func NewShopService(repo ShopRepo, log logger.Logger) *ShopService { // 新增log参数
	return &ShopService{
		repo: repo,
		log:  log, // 初始化日志字段
	}
}
