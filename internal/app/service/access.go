package service

import "loan-admin/internal/app/dto"

type AccessRepo interface {
}

type AccessService struct {
	repo AccessRepo
}

func NewAccessService(repo AccessRepo) *AccessService {
	return &AccessService{repo: repo}
}

func (s *AccessService) GetAccess(id int) (dto.GetUserResp, error) {
	return dto.GetUserResp{}, nil
}
