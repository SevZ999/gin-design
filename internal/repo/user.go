package repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepo struct {
	db  *gorm.DB
	rbd *redis.Client
}

func NewUserRepo(db *gorm.DB, rbd *redis.Client) *UserRepo {
	return &UserRepo{
		db:  db,
		rbd: rbd,
	}
}
