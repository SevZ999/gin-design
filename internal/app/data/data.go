package data

import (
	"context"

	"gorm.io/gorm"
)

// 定义自定义键类型
type txKeyType struct{}

var (
	txKey    txKeyType
	globalDb *gorm.DB
)

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
	//断言是否有事务传入
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok {
		return tx.WithContext(ctx)
	} else {
		return d.db.WithContext(ctx)
	}
}
