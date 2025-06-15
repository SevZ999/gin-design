package repo

import (
	"context"
	"gin-design/internal/app/data"
	"gin-design/internal/app/model"
)

type UserRepo struct {
	Db *data.Data
}

func NewUserRepo(db *data.Data) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (u *UserRepo) GetUser(ctx context.Context, id int) (model.User, error) {

	// err1 := u.Db.WithContext(ctx).Clauses(
	// 	clause.OnConflict{
	// 		Columns:   []clause.Column{{Name: "id"}},
	// 		DoUpdates: clause.AssignmentColumns([]string{"click_time"}),
	// 	},
	// ).Create(&model.User{}).Error

	var user model.User
	err := u.Db.WithContext(ctx).Model(&model.User{}).First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil

}
