//go:build wireinject

package internal

import (
	"loan-admin/internal/app/controller"
	"loan-admin/internal/app/data"
	"loan-admin/internal/app/repo"
	"loan-admin/internal/app/router"
	"loan-admin/internal/app/service"
	"loan-admin/internal/config"
	"loan-admin/internal/db"
	"loan-admin/internal/middleware"
	"loan-admin/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "loan-admin/docs"
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

	api.Use(middleware.CORS())

	a.SetRoute(api)

	a.Engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// pprof.Register(a.Engine, "/api/pprof")

	a.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	a.Engine.Run(":9001")
}

func InitApp(mode string) (*App, error) {

	wire.Build(
		config.LoadConfig,

		logger.NewZapLogger,
		db.NewGormDB,

		data.NewData,
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
