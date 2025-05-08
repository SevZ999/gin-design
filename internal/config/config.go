package config

import (
	"gin-design/pkg/log"

	"github.com/gin-gonic/gin"
)

type Loader interface {
	Load()
}

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() *gin.Engine {
	//TODO 初始化配置和一些全局实例

	loading([]Loader{
		log.NewLogger(),
	})

	return gin.Default()
}

func loading(loader []Loader) {
	for _, l := range loader {
		l.Load()
	}
}
