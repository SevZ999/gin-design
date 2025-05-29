package service

import (
	"context"
	"errors"
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/model"
	"loan-admin/internal/pkg/auth"
	"loan-admin/internal/pkg/logger"
	"time"

	"go.uber.org/zap"
)

//go:generate mockgen -package mock -destination mock_user.go -source user.go
type UserRepo interface {
	GetUser(tx context.Context, id int) (model.User, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (dto.GetUserResp, error) {

	user, err := s.repo.GetUser(context.Background(), id)
	if err != nil {
		return dto.GetUserResp{}, err
	}

	return dto.GetUserResp{
		Id:   user.Id,
		Name: user.Phone,
	}, err
}

func (s *UserService) Login(req dto.LoginReq) (*dto.LoginResp, error) {

	if req.Username != "" && req.Password != "" {
		user, err := s.repo.GetUser(context.Background(), 201)
		if err != nil {
			logger.Log().Error(err.Error())
			return nil, err
		}

		token, err := auth.GenerateToken("secret", "201", user.Name, user.Phone, 24*time.Hour)
		if err != nil {
			zap.L().Sugar().Error(err)
			return nil, err
		}
		return &dto.LoginResp{
			Token: token,
		}, nil

	}

	logger.Log().Error("invalid username or password")

	return nil, errors.New("invalid username or password")
}
