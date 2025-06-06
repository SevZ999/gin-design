//go:build wireinject

package internal

import (
	"gin-design/internal/app/controller"
	"gin-design/internal/app/data"
	"gin-design/internal/app/repo"
	"gin-design/internal/app/router"
	"gin-design/internal/app/service"
	"gin-design/internal/config"
	"gin-design/internal/db"
	"gin-design/internal/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"fmt"
	_ "gin-design/docs"
)

type App struct {
	srv *http.Server
}

func NewApp(cfg *config.Config, handler *gin.Engine) *App {
	return &App{
		srv: &http.Server{
			Handler: handler,
			Addr:    fmt.Sprint(":", cfg.HTTP.Port),
		},
	}
}

func (a *App) Run() {
	if err := a.srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func InitApp(mode string) (*App, func(), error) {

	wire.Build(
		config.LoadConfig,

		logger.NewZapLogger,
		wire.Bind(new(logger.Logger), new(*logger.ZapLogger)),

		db.NewGormDB,

		data.NewData,
		repo.RepoProviderSet,
		service.ServiceProviderSet,
		controller.ControllerProviderSet,
		router.RouterProviderSet,
		router.NewRouters,
		NewEngine,
		NewApp,
	)
	return nil, nil, nil
}
