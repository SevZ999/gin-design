package repo

import (
	"loan-admin/internal/app/service"

	"github.com/google/wire"
)

var RepoProviderSet = wire.NewSet(
	NewUserRepo,
	wire.Bind(new(service.UserRepo), new(*UserRepo)),

	NewAccessRepo,
	wire.Bind(new(service.AccessRepo), new(*AcessRepo)),

	NewChannelRepo,
	wire.Bind(new(service.ChannelRepo), new(*ChannelRepo)),

	NewShopRepo,
	wire.Bind(new(service.ShopRepo), new(*ShopRepo)),
)
