#!/bin/bash

# 检查参数是否存在
if [ $# -eq 0 ]; then
    echo "使用方法：$0 <模块名（小写）>"
    exit 1
fi

MODULE_NAME=$(echo "$1" | sed 's/^\(.\)/\U\1/')
MODULE_LOWER=$1

declare -A FILES=(
    ["controller"]="package controller

type ${MODULE_NAME}Controller struct {
	srv *service.${MODULE_NAME}Service
}

func New${MODULE_NAME}Controller(userService *service.${MODULE_NAME}Service) *${MODULE_NAME}Controller {
	return &${MODULE_NAME}Controller{
		srv: userService,
	}
}"

    ["service"]="package service

type ${MODULE_NAME}Repo interface {
}

type ${MODULE_NAME}Service struct {
	repo ${MODULE_NAME}Repo
	log  *logger.Logger  // 新增日志字段
}

func New${MODULE_NAME}Service(repo ${MODULE_NAME}Repo, log *logger.Logger) *${MODULE_NAME}Service {  // 新增log参数
	return &${MODULE_NAME}Service{
		repo: repo,
		log:  log,  // 初始化日志字段
	}
}"

    ["repo"]="package repo

type ${MODULE_NAME}Repo struct {
	Db *data.Data
}

func New${MODULE_NAME}Repo(db *data.Data) *${MODULE_NAME}Repo {
	return &${MODULE_NAME}Repo{
		Db: db,
	}
}"

    ["router"]="package router

type ${MODULE_NAME}Router struct {
	ctrl *controller.${MODULE_NAME}Controller
}

func New${MODULE_NAME}Router(ctrl *controller.${MODULE_NAME}Controller) *${MODULE_NAME}Router {
	return &${MODULE_NAME}Router{
		ctrl: ctrl,
	}
}

func (r *${MODULE_NAME}Router) SetRoute(router *gin.RouterGroup) {
}"
)

for dir in "${!FILES[@]}"; do
    mkdir -p "./internal/app/$dir"
    filename="./internal/app/$dir/${MODULE_LOWER}.go"
    
    if [ -f "$filename" ]; then
        echo "文件已存在，跳过：$filename"
        continue
    fi
    
    echo "正在生成：$filename"
    
    # 使用here document写入内容（自动处理缩进）
    cat > "$filename" <<EOF
${FILES[$dir]}
EOF
done

echo "文件生成完成！"