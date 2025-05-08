package main

import (
	"gin-design/internal/app"
	"gin-design/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg.Init())

	app.Run()
}
