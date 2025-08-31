# 阶段1：Go程序构建
FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o marktable ./cmd/main.go

# 阶段2：最终镜像
FROM alpine:latest

# 使用国内镜像源替换默认源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装LibreOffice和必要依赖
RUN apk update && apk add \
    libreoffice \
    libreoffice-writer \
    openjdk17-jre \
    ttf-dejavu \
    ttf-liberation \
    font-noto-cjk \
    bash \
    curl \
    && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/marktable .
COPY ./config/config.yaml ./config/

ENV PATH="/usr/lib/libreoffice/program:${PATH}"
EXPOSE 8080
CMD ["./marktable"]