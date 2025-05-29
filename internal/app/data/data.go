package data

import (
	"context"

	"gorm.io/gorm"
)

var globalDb *gorm.DB

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) *Data {

	globalDb = db

	return &Data{
		db: db,
	}
}

func GetDB() *gorm.DB {
	return globalDb
}

func (d *Data) WithContext(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return d.db
	}

	tx := ctx.Value("tx")
	if tx != nil {
		return tx.(*gorm.DB)
	}
	return d.db
}
