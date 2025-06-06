package service

import (
	"gin-design/internal/app/dto"
	"gin-design/internal/app/model"
	"gin-design/internal/pkg/logger"

	"go.uber.org/zap"
)

type UserRepo interface {
	GetUser(int) (model.User, error)
}

type UserService struct {
	repo UserRepo
	log  logger.Logger
}

func (u *UserService) GetUser(req dto.GetUserReq) (dto.GetUserResp, error) {
	user, err := u.repo.GetUser(200)
	if err != nil {
		u.log.Error("get user error", zap.Error(err))
		return dto.GetUserResp{}, err
	}
	return dto.GetUserResp{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func NewUserService(repo UserRepo, log logger.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}
