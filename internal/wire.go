//go:build wireinject
// +build wireinject

package internal

import (
	"gin-design/internal/app/controller"
	"gin-design/internal/app/repo"
	"gin-design/internal/app/router"
	"gin-design/internal/app/service"
	"gin-design/internal/config"
	"gin-design/internal/db"
	"gin-design/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	Engine  *gin.Engine
	routers []router.Router
}

func NewApp(routers []router.Router) *App {
	return &App{
		Engine:  gin.Default(),
		routers: routers,
	}
}

func (a *App) Run() {

	api := gin.Default().Group("/api")

	a.SetRoute(api)

	a.Engine.Run(":9001")
}

func InitApp(mode string) (*App, error) {

	wire.Build(
		config.LoadConfig,

		logger.NewZapLogger,
		db.NewGormDB,

		repo.RepoProviderSet,
		service.ServiceProviderSet,
		controller.ControllerProviderSet,
		router.RouterProviderSet,
		router.NewRouters,
		NewApp,
	)
	return &App{}, nil
}

func (a *App) SetRoute(api *gin.RouterGroup) {
	for _, r := range a.routers {
		r.SetRoute(api)
	}
}
