package service

import "gin-design/internal/app/dto"

type UserRepo interface {
	GetUser(id int) (dto.GetUserResp, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (dto.GetUserResp, error) {
	return s.repo.GetUser(id)
}
