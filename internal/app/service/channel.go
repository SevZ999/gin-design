package service

type ChannelRepo interface {
}

type ChannelService struct {
	ChannelRepo ChannelRepo
}

func NewChannelService(repo ChannelRepo) *ChannelService {
	return &ChannelService{
		ChannelRepo: repo,
	}
}
