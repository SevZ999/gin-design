#!/bin/bash

# 检查参数是否存在
if [ $# -eq 0 ]; then
    echo "使用方法：$0 <模块名（小写）>"
    exit 1
fi

# 将输入转换为首字母大写格式（如 user → User）
MODULE_NAME=$(echo "$1" | sed 's/^\(.\)/\U\1/')
MODULE_LOWER=$1

# 定义要生成的文件结构
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
}

func New${MODULE_NAME}Service(repo ${MODULE_NAME}Repo) *${MODULE_NAME}Service {
	return &${MODULE_NAME}Service{repo: repo}
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

# 创建目录并生成文件
for dir in "${!FILES[@]}"; do
    # 创建目录（如果不存在）
    mkdir -p "./internal/app/$dir"
    
    # 生成文件路径
    filename="./internal/app/$dir/${MODULE_LOWER}.go"
    
    # 检查文件是否存在
    if [ -f "$filename" ]; then
        echo "文件已存在，跳过：$filename"
        continue
    fi
    
    echo "正在生成：$filename"
    
    # 写入文件内容（自动格式化）
    cat > "$filename" <<EOF
${FILES[$dir]}
EOF
done

echo "文件生成完成！"