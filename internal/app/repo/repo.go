package repo

import (
	"gin-design/internal/app/service"

	"github.com/google/wire"
)

var RepoProviderSet = wire.NewSet(
	NewUserRepo,
	wire.Bind(new(service.UserRepo), new(*UserRepo)),

	NewShopRepo,
	wire.Bind(new(service.ShopRepo), new(*ShopRepo)),
)
