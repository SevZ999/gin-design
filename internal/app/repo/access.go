package repo

import (
	"loan-admin/internal/app/data"
	"loan-admin/internal/app/dto"
)

type AcessRepo struct {
	Db *data.Data
}

func NewAccessRepo(db *data.Data) *AcessRepo {
	return &AcessRepo{
		Db: db,
	}
}

func (r *AcessRepo) GetAccess(id int) (dto.GetUserResp, error) {
	return dto.GetUserResp{}, nil
}
