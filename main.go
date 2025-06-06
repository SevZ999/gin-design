package main

import (
	"fmt"
	"gin-design/internal"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		// 开发环境建议警告，生产环境可忽略（无 .env 文件时）
		fmt.Println("Warning: .env file not found")
	}

	// 加载 .env 文件（优先加载）
	name, ok := os.LookupEnv(gin.EnvGinMode)
	if ok {
		gin.SetMode(name)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app, clean, err := internal.InitApp("dev")
	if err != nil {
		panic(err)
	}
	defer clean()

	app.Run()
}
