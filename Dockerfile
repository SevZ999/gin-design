# 第一阶段：构建阶段
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# 设置国内代理加速
ENV GOPROXY=https://goproxy.cn,direct

# 预编译依赖
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# 🔥 全量复制所有文件（注意配合.dockerignore使用）
COPY . .

# 构建二进制文件（静态链接）
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd

# 第二阶段：生产环境镜像
FROM alpine:latest

WORKDIR /app

# 🔥 仅复制必要文件到最终镜像
COPY --from=builder /app/main .
COPY --from=builder /app/config/config.dev.yaml ./config/

# 安装证书包（确保 TLS 连接正常）
# RUN apk add --no-cache ca-certificates

EXPOSE 9001

# 启动命令
CMD ["./main"]