package main

import (
	"fmt"
	"loan-admin/internal"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	app, err := internal.InitApp("dev")
	if err != nil {
		panic(err)
	}

	app.Run()
}
