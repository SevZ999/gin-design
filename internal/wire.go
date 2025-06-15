//go:build wireinject

package internal

import (
	"context"
	"gin-design/internal/app/controller"
	"gin-design/internal/app/data"
	"gin-design/internal/app/repo"
	"gin-design/internal/app/router"
	"gin-design/internal/app/service"
	"gin-design/internal/config"
	"gin-design/internal/db"
	"gin-design/internal/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			Addr:    fmt.Sprint("0.0.0.0:", cfg.HTTP.Port),
		},
	}
}

func (a *App) Run() {
	go func() {
		log.Printf("server started: %s", a.srv.Addr)
		if err := a.srv.ListenAndServe(); err != nil {
			// 仅当错误不是 http.ErrServerClosed 时触发 panic
			if err != http.ErrServerClosed {
				log.Fatalf("服务器意外关闭或启动失败: %v", err)
			}
			// 正常关闭（如调用 Shutdown）时不处理
		}
	}()

	// 监听 SIGTERM 和 SIGINT（Ctrl+C）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// 关闭服务器（停止接收新请求，等待已有请求完成）
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Fatalf("server close error: %v", err)
	}
	log.Println("server stopped")
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
