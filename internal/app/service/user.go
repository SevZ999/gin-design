package service

import (
	"context"
	"gin-design/internal/app/data"
	"gin-design/internal/app/dto"
	"gin-design/internal/app/model"
	"gin-design/internal/pkg/logger"

	"go.uber.org/zap"
)

type UserRepo interface {
	GetUser(context.Context, int) (model.User, error)
}

type UserService struct {
	repo UserRepo
	log  logger.Logger
}

func (u *UserService) GetUser(ctx context.Context, req dto.GetUserReq) (dto.GetUserResp, error) {

	var resp dto.GetUserResp

	tx := data.NewTransactionManager(ctx)

	err := tx.ExecuteTransaction(func(txCtx context.Context) error {
		user, err := u.repo.GetUser(txCtx, req.Id)
		if err != nil {
			return err
		}
		resp = dto.GetUserResp{
			Id:   user.Id,
			Name: user.Name,
		}
		return nil
	})

	if err != nil {
		u.log.Error(ctx, "get user error", zap.Error(err))
		return dto.GetUserResp{}, err
	}

	return resp, nil
}

func NewUserService(repo UserRepo, log logger.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}
