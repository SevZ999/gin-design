package router

import "github.com/gin-gonic/gin"

type Routers interface {
	Route(*gin.RouterGroup)
}

type router struct {
	engine *gin.Engine
}

func NewRouter() *router {
	return &router{}
}

func (r *router) Registe() {

	// 初始化前缀
	api := r.engine.Group("/api")

	// 全局中间件
	r.engine.Use(gin.Recovery())

	// 注册路由
	routers := []Routers{
		NewUserRouter(nil),
	}

	for _, router := range routers {
		router.Route(api)
	}
}
