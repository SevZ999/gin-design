package repo

import (
	"context"
	"loan-admin/internal/app/data"
	"loan-admin/internal/app/model"
)

type UserRepo struct {
	Db *data.Data
}

func NewUserRepo(db *data.Data) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (r *UserRepo) GetUser(tx context.Context, id int) (model.User, error) {

	user := model.User{}
	err := r.Db.WithContext(tx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByName(tx context.Context, username string) (model.User, error) {

	user := model.User{}
	err := r.Db.WithContext(tx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
