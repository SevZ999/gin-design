package repo

import "loan-admin/internal/app/data"

type ChannelRepo struct {
	Db *data.Data
}

func NewChannelRepo(db *data.Data) *ChannelRepo {
	return &ChannelRepo{Db: db}
}
