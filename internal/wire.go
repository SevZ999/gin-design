//go:build wireinject
// +build wireinject

package internal

import (
	"loan-admin/internal/app/controller"
	"loan-admin/internal/app/repo"
	"loan-admin/internal/app/router"
	"loan-admin/internal/app/service"
	"loan-admin/internal/config"
	"loan-admin/internal/db"
	"loan-admin/internal/pkg/logger"

	"github.com/gin-contrib/pprof"
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

	api := a.Engine.Group("api")

	a.SetRoute(api)

	a.Engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pprof.Register(a.Engine, "/api/pprof")

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
