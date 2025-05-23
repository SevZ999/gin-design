package service

import (
	"errors"
	"loan-admin/internal/app/dto"
	"loan-admin/internal/app/model"
	"loan-admin/internal/pkg/auth"
	"time"
)

type UserRepo interface {
	GetUser(id int) (model.User, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (dto.GetUserResp, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return dto.GetUserResp{}, err
	}

	return dto.GetUserResp{
		Id:   user.Id,
		Name: user.Name,
	}, err
}

func (s *UserService) Login(req dto.LoginReq) (*dto.LoginResp, error) {
	if req.Username == "admin" && req.Password == "admin" {
		token, err := auth.GenerateToken("secret", "1", "test", "admin", 24*time.Hour)
		if err != nil {
			return nil, err
		}
		return &dto.LoginResp{
			Token: token,
		}, nil
	}

	return nil, errors.New("invalid username or password")
}
