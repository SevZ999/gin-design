package main

import (
	"context"
	"gin-design/internal"
	"gin-design/internal/pkg/logger"
	"os"
)

// @title 后台管理系统
// @version 1.0
// @description 后台管理系统
// @host https://api.example.com
// @BasePath /api
// @schemes https
// @accept json
// @produce json
func main() {

	// 加载 .env 文件（优先加载）
	name, ok := os.LookupEnv("MODE")
	if !ok {
		name = "debug"
	}

	app, clean, err := internal.InitApp(name)
	if err != nil {
		logger.GetLogger().Fatal(context.TODO(), err.Error())
	}
	defer clean()

	app.Run()

}
