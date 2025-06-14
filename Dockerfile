# 第一阶段：构建阶段
FROM golang:1.24.2-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct

# 预下载依赖（利用Docker缓存层）
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并构建
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main ./cmd

# 第二阶段：生产镜像
FROM alpine:latest

# 创建非root用户和组，安装CA证书
RUN addgroup -g 1001 nonroot && \
    adduser -u 1001 -G nonroot -D nonroot && \
    apk add --no-cache ca-certificates

WORKDIR /app

# 复制二进制文件
COPY --from=builder --chown=nonroot:nonroot /app/main .

# 创建配置目录（不复制配置文件）
RUN mkdir -p ./config && chmod 750 ./config

# 设置容器用户和启动命令
USER nonroot:nonroot
CMD ["./main"]
