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

# 复制二进制文件和配置文件（生产环境专用）
COPY --from=builder --chown=nonroot:nonroot /app/main .
COPY --from=builder --chown=nonroot:nonroot /app/config/config.dev.yaml ./config/

# 确保配置文件权限安全
RUN chmod 640 ./config/config.dev.yaml

# 设置容器用户和启动命令
USER nonroot:nonroot
CMD ["./main"]