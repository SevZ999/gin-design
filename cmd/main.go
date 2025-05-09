package main

import (
	"gin-design/internal/app"
	"gin-design/internal/config"

	"github.com/google/wire"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg.Init())

	wire.StructProvider

	app.Run()
}
