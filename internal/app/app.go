package app

import (
	"gin-design/internal/app/router"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
}

func NewApp(engine *gin.Engine) *App {
	return &App{
		Engine: engine,
	}
}

func (a *App) Run() {
	r := router.NewRouter()
	r.Registe()
	a.Engine.Run()
}
