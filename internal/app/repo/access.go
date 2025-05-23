package repo

import (
	"loan-admin/internal/app/dto"

	"gorm.io/gorm"
)

type AcessRepo struct {
	db *gorm.DB
}

func NewAccessRepo(db *gorm.DB) *AcessRepo {
	return &AcessRepo{
		db: db,
	}
}

func (r *AcessRepo) GetAccess(id int) (dto.GetUserResp, error) {
	return dto.GetUserResp{}, nil
}
