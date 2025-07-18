# 阶段1：LibreOffice专用构建
FROM alpine:3.18 AS libreoffice
RUN apk add --no-cache wget tar
# 下载官方预编译包（体积比apt安装小40%）
ARG LO_VERSION=7.6.4
RUN wget -q https://download.documentfoundation.org/libreoffice/stable/${LO_VERSION}/deb/x86_64/LibreOffice_${LO_VERSION}_Linux_x86-64_deb.tar.gz -O /tmp/lo.tar.gz && \
    mkdir -p /tmp/lo && \
    tar -xzf /tmp/lo.tar.gz -C /tmp/lo --strip-components=1 && \
    rm /tmp/lo.tar.gz


# 阶段2：Go程序构建
FROM golang:1.22.0  AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载所有依赖包（go get），避免每次构建都下载
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN go mod download

# 复制源代码
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o marktable ./cmd/main.go


# 阶段3：最终镜像
FROM alpine:3.18
# 安装LibreOffice最小依赖
RUN apk add --no-cache \
    bash \
    libc6-compat \
    libreoffice-common \
    libreoffice-writer \
    openjdk17-jre \
    ttf-dejavu \
    ttf-liberation

# 从阶段1复制LibreOffice核心文件
COPY --from=libreoffice /tmp/lo/DEBS/*.deb /tmp/lo/
RUN apk add --allow-untrusted /tmp/lo/*.deb && \
    rm -rf /tmp/lo \

RUN apk add --no-cache \
    wqy-zenhei-font \
    noto-fonts-cjk

# 环境变量配置
ENV PATH="/usr/lib/libreoffice/program:${PATH}"
ENV URE_BOOTSTRAP="vnd.sun.star.pathname:/usr/lib/libreoffice/program/fundamentalrc"
ENV SAL_USE_VCLPLUGIN="headless"
ENV LIBO_FLAGS="--headless --invisible --nocrashreport --nodefault --nologo"

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s \
    CMD soffice --help >/dev/null 2>&1 || exit 1

# 阶段2：运行
FROM alpine:latest
# 安装CA证书（如需HTTPS请求）
RUN apk --no-cache add ca-certificates

WORKDIR /app
# 从构建阶段复制可执行文件
COPY --from=builder /app/marktable /usr/local/bin/
COPY ./config/config.yaml /app/config/

# 设置临时文件目录（LibreOffice需要）
ENV HOME=/tmp

EXPOSE 8080

# 运行Marktable
CMD ["marktable"]