package repo

import (
	"gin-design/internal/app/dto"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUser(id int) (dto.GetUserResp, error) {
	var user dto.GetUserResp
	err := r.db.Table("users").Where("id = ?", id).First(&user).Error
	return user, err
}
