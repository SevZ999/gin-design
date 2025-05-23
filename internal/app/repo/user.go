package repo

import (
	"loan-admin/internal/app/model"

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

func (r *UserRepo) GetUser(id int) (model.User, error) {
	return model.User{
		Id:   id,
		Name: "test",
	}, nil
}
