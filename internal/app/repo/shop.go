package repo

import "loan-admin/internal/app/data"

type ShopRepo struct {
	Db *data.Data
}

func NewShopRepo(db *data.Data) *ShopRepo {
	return &ShopRepo{
		Db: db,
	}
}
