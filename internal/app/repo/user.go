package repo

import "loan-admin/internal/app/data"

type UserRepo struct {
	Db *data.Data
}

func NewUserRepo(db *data.Data) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}
