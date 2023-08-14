# 使用官方Golang镜像作为基础镜像
FROM golang:latest as builder

# 在容器中创建应用目录
WORKDIR /app

# 将主机 `gin` 目录下的内容复制到容器的 `/app` 目录下
COPY . .

# 开始构建 Go 应用
RUN go mod download
RUN GOOS=linux GOARCH=arm go build -o webook .

# 暴露端口，允许外部连接到 Go 应用（需要根据你的 Go 应用配置来调整）
EXPOSE 8080


# 设置启动命令
ENTRYPOINT [ "./webook" ]