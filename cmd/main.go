package main

import (
	"fmt"
	"gin-design/internal"

	"github.com/joho/godotenv"
)

func main() {

	// 加载 .env 文件（优先加载）
	if err := godotenv.Load(); err != nil {
		// 开发环境建议警告，生产环境可忽略（无 .env 文件时）
		fmt.Println("Warning: .env file not found")
	}

	app, err := internal.InitApp("dev")
	if err != nil {
		panic(err)
	}

	asd := make(map[string]string)
	asd["ads"] = "ads"

	app.Run()
}
